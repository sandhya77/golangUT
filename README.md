OpenEBS

http://www.openebs.io/

OpenEBS is purpose built storage for containerized environments, written in GoLang, built upon block storage containers we call Virtual Storage Machines or VSMs.

#golangUT
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. 
A service provided to install go with the help of vagrantfile.

INSTALLATION AND GETTING STARTED:

TheRe are some of the steps to start with the development and running of go in one's laptop.
It requires LINUX OS (ubuntu16.04).
In addition the laptop has virtualbox(5.0) and vagrant(1.9.3) to be installed.

- git clone https://github.com/sandhya77/golangUT.git 
- cd to above cloned folder i.e golangUT
- vagrant up
  (The initial setup is, depending on your internet connection, going to take some time. It's normal for VirtualBox to show up, please don't login and wait for the installation to finish).
 - vagrant ssh
 - Inside the vagrant VM repeat the above steps again
 - check with goenv
 - go test
 
TROUBLE SHOOTING DURING LOCAL SETUP:
 
 - 'vagrant up' is a time taking process
 - This downloads all the vendoring libraries from vagrantfile
 - Typically time requirement is more for the first attempt only
 - It may show network issues as errors and sometimes errors due to "Intell       Virtualization Configuration" to be enabled in ur laptop.
