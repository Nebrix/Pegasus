using System.Diagnostics;
using System.Net;
using System.Net.NetworkInformation;
using System.Net.Sockets;
using System.Text.Json;

namespace PScan
{
    class Program
    {
        private const string Uri = "https://pscan.com";
        private static bool ProceedIfHostDown;

        static async Task Main(string[] args)
        {
            if (args.Length < 1 || args[0] == "--help")
            {
                PrintUsage();
                return;
            }

            string target = args[0];
            bool isIPAddress = IsIPAddress(target);

            string ipAddress = isIPAddress ? target : ResolveIPv4Address(target);
            string hostName = isIPAddress ? ResolveHostName(ipAddress) : target;

            ProceedIfHostDown = args.Contains("-Pn");

            if (!ProceedIfHostDown && !IsHostUp(ipAddress))
            {
                Console.WriteLine($"Host network is down or blocking ICMP (ping). Use the -Pn flag to scan the network.");
                return;
            }

            List<int> portsToScan = ParsePortsFromArgs(args);

            if (portsToScan.Count == 0)
            {
                // No specific ports provided, scan all ports
                portsToScan.AddRange(Enumerable.Range(1, 65535));
            }

            Console.WriteLine($"Starting Pscan 0.0.1 ({Uri}) at {GetFormattedTimestamp()}");

            var tasks = new List<Task<PortStatus>>();

            // Start a stopwatch to measure elapsed time
            var stopwatch = new Stopwatch();
            stopwatch.Start();

            foreach (int port in portsToScan)
            {
                tasks.Add(ScanPortAsync(ipAddress, port));
            }

            var results = await Task.WhenAll(tasks);

            Console.WriteLine($"Scan report for {hostName} ({ipAddress})");
            Console.WriteLine($"Host is up ({stopwatch.Elapsed.TotalSeconds:F4}s latency).");
            Console.WriteLine("PORT        STATE    SERVICE");

            foreach (var (port, status) in portsToScan.Zip(results, (p, s) => (p, s)))
            {
                string serviceName = GetServiceName(port);
                string protocol = GetProtocol(port);

                if (status == PortStatus.Open)
                {
                    Console.WriteLine($"{port,5:G}/{protocol.ToLower(),-4}  open     {serviceName}");
                }
            }

            // Stop the stopwatch and print the elapsed time
            stopwatch.Stop();
            Console.WriteLine($"Pscan done: scanned in {stopwatch.Elapsed.TotalSeconds:F2} seconds");
        }

        private static List<int> ParsePortsFromArgs(string[] args)
        {
            var portsToScan = new List<int>();

            for (int i = 1; i < args.Length; i++)
            {
                if (args[i] == "--ports" && i + 1 < args.Length)
                {
                    string[] portStrings = args[i + 1].Split(',');
                    foreach (var portString in portStrings)
                    {
                        if (int.TryParse(portString, out int port) && port >= 1 && port <= 65535)
                        {
                            portsToScan.Add(port);
                        }
                        else
                        {
                            Console.WriteLine($"Invalid port: {portString}");
                        }
                    }
                    i++;
                }
                else if (args[i] == "-Pn")
                {
                    ProceedIfHostDown = true;
                }
                else
                {
                    Console.WriteLine($"Invalid option: {args[i]}");
                }
            }
            return portsToScan;
        }

        private static string GetFormattedTimestamp()
        {
            // Get the current time
            DateTime currentTime = DateTime.Now;

            // Get the timezone abbreviation for the local time zone
            string timeZoneAbbreviation = GetAbbreviationFromTimeZoneName(TimeZoneInfo.Local);

            // Format the timestamp as "yyyy-MM-dd HH:mm TIMEZONE"
            string formattedTimestamp = currentTime.ToString("yyyy-MM-dd HH:mm ") + timeZoneAbbreviation;

            return formattedTimestamp;
        }

        private static string GetAbbreviationFromTimeZoneName(TimeZoneInfo timeZone)
        {
            string timeZoneName = timeZone.IsDaylightSavingTime(DateTime.Now)
                ? timeZone.DaylightName
                : timeZone.StandardName;

            // Extract uppercase letters from the timezone name
            string abbreviation = string.Concat(timeZoneName.Where(char.IsUpper));

            return abbreviation;
        }

        private static void PrintUsage()
        {
            Console.WriteLine("Usage: PScan [options] target");
            Console.WriteLine("Options:");
            Console.WriteLine("  --help             Show this help message");
            Console.WriteLine("  -Pn                Treat all hosts as online -- skip host discovery");
            Console.WriteLine("  --ports <port(s)>  Specify ports to scan (comma-separated)");
        }

        static bool IsIPAddress(string input)
        {
            return IPAddress.TryParse(input, out _);
        }

        static string ResolveIPv4Address(string host)
        {
            try
            {
                IPAddress[] addresses = Dns.GetHostAddresses(host);
                foreach (var address in addresses)
                {
                    if (address.AddressFamily == AddressFamily.InterNetwork)
                    {
                        return address.ToString();
                    }
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Error resolving IP address for {host}: {ex.Message}");
            }

            return "Unknown";
        }

        static bool IsHostUp(string host)
        {
            try
            {
                using Ping ping = new();
                PingReply reply = ping.Send(host);
                return reply.Status == IPStatus.Success;
            }
            catch (PingException ex)
            {
                Console.WriteLine($"Error checking host status for {host}: {ex.Message}");
                return false;
            }
        }

        static string ResolveHostName(string ip)
        {
            try
            {
                IPHostEntry hostEntry = Dns.GetHostEntry(IPAddress.Parse(ip).MapToIPv4());
                return hostEntry.HostName;
            }
            catch (SocketException)
            {
                return "No Host Name Found";
            }
            catch (Exception ex)
            {
                return "Error: " + ex.Message;
            }
        }

        static async Task<PortStatus> ScanPortAsync(string host, int port)
        {
            using var client = new TcpClient();
            try
            {
                var connectTask = client.ConnectAsync(host, port);
                var timeoutTask = Task.Delay(2000);

                var completedTask = await Task.WhenAny(connectTask, timeoutTask);

                if (completedTask == connectTask)
                {
                    return PortStatus.Open;
                }
                else
                {
                    return PortStatus.Filtered;
                }
            }
            catch (SocketException)
            {
                return PortStatus.Filtered;
            }
            catch (Exception)
            {
                return PortStatus.Filtered;
            }
        }

        private static List<PortInfo> cachedPortInfos = null!;

        static string GetServiceName(int port)
        {
            if (cachedPortInfos == null)
            {
                string fileName = "ports.json";
                try
                {
                    cachedPortInfos = ReadPortInfoFromFile(fileName);
                }
                catch (IOException ex)
                {
                    Console.WriteLine(ex.Message);
                    cachedPortInfos = new List<PortInfo>();
                }
            }

            PortInfo portInfo = cachedPortInfos.FirstOrDefault(p => p.Port == port)!;
            return portInfo?.Service ?? "unknown";
        }

        static string GetProtocol(int port)
        {
            if (port >= 0 && port <= 1023)
            {
                return "TCP";
            }
            else if (port >= 1024 && port <= 49151)
            {
                return "UDP";
            }
            else
            {
                return "Unknown";
            }
        }

        static List<PortInfo> ReadPortInfoFromFile(string fileName)
        {
            List<PortInfo> portInfos = new List<PortInfo>();

            try
            {
                string jsonContent = File.ReadAllText(fileName);

                var jsonObj = JsonSerializer.Deserialize<Dictionary<string, string>>(jsonContent);

                if (jsonObj != null)
                {
                    foreach (var kvp in jsonObj)
                    {
                        if (int.TryParse(kvp.Key, out int port))
                        {
                            PortInfo portInfo = new PortInfo
                            {
                                Port = port,
                                Service = kvp.Value
                            };

                            portInfos.Add(portInfo);
                        }
                    }
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
            }

            return portInfos;
        }

        enum PortStatus
        {
            Open,
            Filtered
        }

        public class PortInfo
        {
            public int Port { get; set; }
            public string? Service { get; set; }
        }
    }
}