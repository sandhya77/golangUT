
INSTALLATION AND GETTING STARTED WITH GO:
There are some steps to be followed in setting up GO envirournment.
REQUIREMENTS:
- Linux os(latest version ubuntu16.04).
- Virtualbox(>=5.0)
- Vagrant(>=1.9.3)
- git
GO INSTALLATION:
- $ git clone https://github.com/sandhya77/golangUT.git 
- $ cd to above cloned folder i.e golangUT
- $ vagrant up
  (The initial setup is, depending on your internet connection, going,takes some time. It's normal for VirtualBox to show up, please don't login and wait for the installation to finish).
 -$ vagrant ssh
 After succesful vagrant up,
 - check with go env
 - go test
 
TROUBLE SHOOTING DURING LOCAL SETUP:
 
 - 'vagrant up' is a time taking process
 - This downloads all the vendoring libraries from vagrantfile
 - Typically time requirement is more for the first attempt only
 - It may show network issues as errors and sometimes errors due to "Intell       Virtualization Configuration" to be enabled in ur laptop.
