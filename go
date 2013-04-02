#!/bin/bash


die(){
  echo $1
  exit 1
}

uid=`id -u`
if [ $uid != 0 ]; then
	die "Usage: sudo $0 [true | false] (Accept/Decline Oracle License Agreement) [username] [password] (Neo4J username/password)"
fi

[ "${1}" == "true" ] && oracle=true || oracle=false
username=$2
password=$3

set_distro(){
	result=`command -v apt-get > /dev/null 2>&1`
	if [ $? == 0  ];then
		echo 'ubuntu'
	fi
}


oracle_license(){
	    if [ "${oracle}" == true ]; then
		echo "********************************************************************************"
		echo "***** By using this you agree to the terms of the Oracle License Agreement *****"
		echo "********************************************************************************"
		export I_ACCEPT_THE_ORACLE_LICENSE=true
	    fi
}

download_neo4j_repo(){
	    wget -q -O neo4j-puppet.zip https://github.com/acquaintable/neo4j-puppet/archive/master.zip
}

install_deb_puppet(){
	wget -q http://apt.puppetlabs.com/puppetlabs-release-$1.deb
	dpkg -i puppetlabs-release-$1.deb
}

ubuntu_install(){
	cd /var/tmp
	distro_version=`lsb_release -c | awk '{print $2}'`
	case $distro_version in
	'squeeze')
		install_deb_puppet $distro_version
		;;
	'lenny')
		install_deb_puppet $distro_version
		;;
	'wheezy')
		install_deb_puppet $distro_version
		;;
	'sid')
		install_deb_puppet $distro_version
		;;
	'precise')
		install_deb_puppet $distro_version
		;;
	'lucid')
		install_deb_puppet $distro_version
		;;
	'hardy')
		install_deb_puppet $distro_version
		;;
	'quantal')
		install_deb_puppet $distro_version
		;;
	'oneiric')
		install_deb_puppet $distro_version
		;;
	'natty')
		install_deb_puppet $distro_version
		;;
	'maverick')
		install_deb_puppet $distro_version
		;;

	*)
		die "$distro_version Ubuntu/Debian distro unsupported by Puppet at this time."
		;;
	esac

	apt-get update -qq
	apt-get install unzip puppet -y
	(
	    cd /var/tmp
	    download_neo4j_repo
	    unzip -q -o neo4j-puppet.zip
	    mv -f neo4j-puppet-master neo4j
	    oracle_license
	    export NEO4J_USERNAME=$username
	    export NEO4J_PASSWORD=$password

	    puppet apply neo4j/tests/init.pp --modulepath . | tee /var/tmp/puppet.log
	)
}


distro=$(set_distro)

case $distro in
	'ubuntu')
		ubuntu_install
		;;
	*)
		die "Only Ubuntu/Debian distributions are supported"
		;;
esac
