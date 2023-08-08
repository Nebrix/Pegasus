use Cwd;

my $cwd = getcwd();

sub string_last_component {
    my $path = shift;
    my $last_dir = (split('/', $path))[-1];
    return $last_dir;
}

my $last_dir = string_last_component($cwd);
print "\033[1;34m$last_dir\033[1;32m ➜\033[0m ";