use strict;
use warnings;
use JSON;
use Cwd;

# Read the configuration from the JSON file
my $config_file = 'src/shell/powerline/config/config.json';
open my $fh, '<', $config_file or die "Could not open $config_file: $!";
my $config_json = do { local $/; <$fh> };
close $fh;

my $config = decode_json($config_json);

# Get the color codes from the configuration
my $blue_color = $config->{colorCodes}{blue};
my $green_color = $config->{colorCodes}{green};

# Get the current working directory and extract the last component
my $cwd = getcwd();
my $last_dir = (split('/', $cwd))[-1];

# Print the last directory with the specified color codes
print "\033[${blue_color}m$last_dir\033[${green_color}m ➜\033[0m ";
