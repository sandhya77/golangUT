VAGRANT – GO INSTALLATION

GETTING STARTED WITH GO:

There are some steps to be followed in setting up GO envirournment.

REQUIREMENTS:

- Linux os on desktop(latest version ubuntu16.04 LTS).
- Virtualbox(>=5.0)
- Vagrant(>=1.9.3)
- git to be installed

GO INSTALLATION ON VM:

- $ git clone https://github.com/sandhya77/golangUT.git 
- $ cd golandUT
- $ vagrant up
  (The initial setup is, depending on your internet connection, going,to take some time).
After succesful vagrant up go with: 
 -$ vagrant ssh
- $ git clone https://github.com/sandhya77/golangUT.git
 - $ go env (1.6)
 -$ go test -coverprofile = ./ (gotest and code coverage)
 
TROUBLE SHOOTING DURING LOCAL SETUP:
 
 - 'vagrant up' is a time taking process,downloads all the vendoring libraries from vagrantfile
 - Typically time requirement is more for the first attempt only
 - It may show network issues as errors and sometimes errors due to "Intell       Virtualization Configuration" to be enabled in ur laptop.
