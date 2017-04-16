#golangUT

A service provided to install go with the help of vagrantfile

SETTING UP GO ENVIROURNMENT AND INSTALLATION

These are some of the steps to start with the development and running of go in one's laptop.

It requires LINUX OS (ubuntu16.04).

In addition the laptop has virtualbox(5.0) and vagrant(1.9.3) to be installed.

 - git clone https://github.com/sandhya77/golangUT.git 
 - cd to above cloned folder i.e golangUT
 - vagrant up
 - vagrant ssh
 
 - Inside the vagrant VM repeat the above steps again
 - check with goenv
 - go test
 
 TROUBLE SHOOTING DURING LOCAL SETUP
 
 -'vagrant up' is a time taking process
 
 -This downloads all the vendoring libraries from vagrantfile
 
 -Typically time requirement is more for the first attempt only
 
 - It may show network issues as errors and sometimes errors due to "Intell Virtualization Configuration" to be enabled in ur       laptop.
