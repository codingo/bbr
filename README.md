# bbr
A simple golang script that will take some template information and process it.

# Arguments
| Argument | Description                      |
|----------|----------------------------------|
| -h       | Display help message and exit.   |
| -t       | Path to template file to use.    |
| -u       | Username to replace _user_ with. |
| -o       | Optional output file name.       |


It will then process the text file, and make the following replacements (not all fields may be present, some will be present more than once):

| Argument      | Description                                               |
|---------------|-----------------------------------------------------------|
| \_target\_      | Replace with the value of the -t argument.                |
| \_username\_    | Replace with the value of the -u argument.                |
| \_sha\_         | Replace with the SHA256 encoded value of the -u argument. |
| \_nameservers\_ | Replace with the output of "dig NS @8.8.8.8 _target_"     |
| \_dig\_         | Replace with the value of "dig @8.8.8.8 _target_"         |
| \_whois\_       | Replace with the whois output of the target parameter.    |
