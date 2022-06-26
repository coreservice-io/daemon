#!/bin/bash

#get system type Linux/FreeBSD/Solaris
SYSTEM=$(uname -s)

#linux_systemd/linux_systemv/linux_upstart

#service name
service_name=$1
if [ $service_name = "" ]; then
  echo "service name error"
  exit
fi

#install|remove|start|stop|restart|status
command=$2

#dir path
project_path=$(
    cd $(dirname $0)
    pwd
)


#linux_systemd
function install_linux_systemd {

    if [ -f /etc/systemd/system/$service_name.service ]; then
        echo "service already exist"
    else

        echo "install..."

        sudo cat >/etc/systemd/system/$service_name.service <<EOF
[Unit]
Description=$service_name
After=network.target

[Service]
StartLimitInterval=15s
StartLimitBurst=5
ExecStart=$project_path/$service_name
StandardOutput=null
StandardError=null
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

        sudo systemctl daemon-reload
        sudo systemctl enable $service_name.service
    fi
}

function remove_linux_systemd {
    stop
    echo "remove..."
    sudo systemctl disable $service_name.service
    sudo rm -f /etc/systemd/system/$service_name.service
    sudo systemctl daemon-reload
}

function start_linux_systemd {
    echo "start..."
    sudo service $service_name start
}

function stop_linux_systemd {
    echo "stop..."
    sudo service $service_name stop
}

function restart_linux_systemd {
    echo "Restarting server.."
    sudo service $service_name restart

}

function status_linux_systemd {
    echo "status"
    sudo service $service_name status
}

#linux_systemv
function install_linux_systemv {

    if [ -f /etc/systemd/system/$service_name.service ]; then
        echo "service already exist"
    else

        echo "install..."

        sudo cat >/etc/systemd/system/$service_name.service <<EOF
[Unit]
Description=$service_name
After=network.target

[Service]
StartLimitInterval=15s
StartLimitBurst=5
ExecStart=$project_path/$service_name
StandardOutput=null
StandardError=null
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

        sudo systemctl daemon-reload
        sudo systemctl enable $service_name.service
    fi
}

function remove_linux_systemv {
    stop
    echo "remove..."
    sudo systemctl disable $service_name.service
    sudo rm -f /etc/systemd/system/$service_name.service
    sudo systemctl daemon-reload
}

function start_linux_systemv {
    echo "start..."
    sudo service $service_name start
}

function stop_linux_systemv {
    echo "stop..."
    sudo service $service_name stop
}

function restart_linux_systemv {
    echo "Restarting server.."
    sudo service $service_name restart

}

function status_linux_systemv {
    echo "status"
    sudo service $service_name status
}


#linux_upstart
function install_linux_upstart {

    if [ -f /etc/systemd/system/$service_name.service ]; then
        echo "service already exist"
    else

        echo "install..."

        sudo cat >/etc/systemd/system/$service_name.service <<EOF
[Unit]
Description=$service_name
After=network.target

[Service]
StartLimitInterval=15s
StartLimitBurst=5
ExecStart=$project_path/$service_name
StandardOutput=null
StandardError=null
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

        sudo systemctl daemon-reload
        sudo systemctl enable $service_name.service
    fi
}

function remove_linux_upstart {
    stop
    echo "remove..."
    sudo systemctl disable $service_name.service
    sudo rm -f /etc/systemd/system/$service_name.service
    sudo systemctl daemon-reload
}

function start_linux_upstart {
    echo "start..."
    sudo service $service_name start
}

function stop_linux_upstart {
    echo "stop..."
    sudo service $service_name stop
}

function restart_linux_upstart {
    echo "Restarting server.."
    sudo service $service_name restart

}

function status_linux_upstart {
    echo "status"
    sudo service $service_name status
}

#linux_freebsd
function install_freebsd {

    if [ -f /etc/systemd/system/$service_name.service ]; then
        echo "service already exist"
    else

        echo "install..."

        sudo cat >/etc/systemd/system/$service_name.service <<EOF
[Unit]
Description=$service_name
After=network.target

[Service]
StartLimitInterval=15s
StartLimitBurst=5
ExecStart=$project_path/$service_name
StandardOutput=null
StandardError=null
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

        sudo systemctl daemon-reload
        sudo systemctl enable $service_name.service
    fi
}

function remove_freebsd {
    stop
    echo "remove..."
    sudo systemctl disable $service_name.service
    sudo rm -f /etc/systemd/system/$service_name.service
    sudo systemctl daemon-reload
}

function start_freebsd {
    echo "start..."
    sudo service $service_name start
}

function stop_freebsd {
    echo "stop..."
    sudo service $service_name stop
}

function restart_freebsd {
    echo "Restarting server.."
    sudo service $service_name restart

}

function status_freebsd {
    echo "status"
    sudo service $service_name status
}


#solaris
function install_solaris {

    if [ -f /etc/systemd/system/$service_name.service ]; then
        echo "service already exist"
    else

        echo "install..."

        sudo cat >/etc/systemd/system/$service_name.service <<EOF
[Unit]
Description=$service_name
After=network.target

[Service]
StartLimitInterval=15s
StartLimitBurst=5
ExecStart=$project_path/$service_name
StandardOutput=null
StandardError=null
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

        sudo systemctl daemon-reload
        sudo systemctl enable $service_name.service
    fi
}

function remove_solaris {
    stop
    echo "remove..."
    sudo systemctl disable $service_name.service
    sudo rm -f /etc/systemd/system/$service_name.service
    sudo systemctl daemon-reload
}

function start_solaris {
    echo "start..."
    sudo service $service_name start
}

function stop_solaris {
    echo "stop..."
    sudo service $service_name stop
}

function restart_solaris {
    echo "Restarting server.."
    sudo service $service_name restart

}

function status_solaris {
    echo "status"
    sudo service $service_name status
}


function getSystemType() {
    if [ $SYSTEM = "Linux" ]; then
      if [[ ! -d "/run/systemd/system" ]]; then
         echo "linux_systemd"
        elif [[ ! -d "/sbin/initctl" ]]; then
         echo "linux_upstart"
        else
          echo "linux_systemv"
        fi
    elif [ $SYSTEM = "FreeBSD" ]; then
      echo "freebsd"
    elif [ $SYSTEM = "Solaris" ]; then
      echo "solaris"
    else
      echo "unsupport"
    fi
}


function test {
  echo $SYSTEM
    echo "project_path:"$project_path
    echo "service_name:"$service_name
    echo "exe_path:"$project_path/$service_name
}

case "$command" in
install)

  case "$getSystemType" in
  linux_systemd)
      install_linux_systemd
      ;;
  linux_upstart)
      install_linux_upstart
      ;;
  linux_systemv)
      install_linux_systemv
      ;;
  freebsd)
      install_freebsd
      ;;
  solaris)
      install_solaris
      ;;
  *)
      echo "unsupport system: $SYSTEM"
      ;;
  esac
    ;;

remove)
    case "$getSystemType" in
      linux_systemd)
          remove_linux_systemd
          ;;
      linux_upstart)
          remove_linux_upstart
          ;;
      linux_systemv)
          remove_linux_systemv
          ;;
      freebsd)
          remove_freebsd
          ;;
      solaris)
          remove_solaris
          ;;
      *)
          echo "unsupport system: $SYSTEM"
          ;;
      esac
    ;;
start)
    case "$getSystemType" in
      linux_systemd)
          start_linux_systemd
          ;;
      linux_upstart)
          start_linux_upstart
          ;;
      linux_systemv)
          start_linux_systemv
          ;;
      freebsd)
          start_freebsd
          ;;
      solaris)
          start_solaris
          ;;
      *)
          echo "unsupport system: $SYSTEM"
          ;;
      esac
    ;;
stop)
    case "$getSystemType" in
      linux_systemd)
          stop_linux_systemd
          ;;
      linux_upstart)
          stop_linux_upstart
          ;;
      linux_systemv)
          stop_linux_systemv
          ;;
      freebsd)
          stop_freebsd
          ;;
      solaris)
          stop_solaris
          ;;
      *)
          echo "unsupport system: $SYSTEM"
          ;;
      esac
    ;;
restart)
    case "$getSystemType" in
      linux_systemd)
          restart_linux_systemd
          ;;
      linux_upstart)
          restart_linux_upstart
          ;;
      linux_systemv)
          restart_linux_systemv
          ;;
      freebsd)
          restart_freebsd
          ;;
      solaris)
          restart_solaris
          ;;
      *)
          echo "unsupport system: $SYSTEM"
          ;;
      esac
    ;;
status)
    case "$getSystemType" in
      linux_systemd)
          status_linux_systemd
          ;;
      linux_upstart)
          status_linux_upstart
          ;;
      linux_systemv)
          status_linux_systemv
          ;;
      freebsd)
          status_freebsd
          ;;
      solaris)
          status_solaris
          ;;
      *)
          echo "unsupport system: $SYSTEM"
          ;;
      esac
    ;;
test)
    test
    ;;
*)
    echo "Usage: sudo $0 {[service_name] install|remove|start|stop|restart|status}"
    ;;
esac
