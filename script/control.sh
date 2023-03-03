#!/bin/bash

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE
cd ../

mkdir -p var &>/dev/null

module=db-go-gin
app=$module
pidfile=var/app.pid
logfile=var/app.log
gitversion=.gitversion

function check_pid() {
    if [ -f $pidfile ];then
        pid=`cat $pidfile`
        if [ -n $pid ]; then
            running=`ps -p $pid|grep -v "PID TTY" |wc -l`
            return $running
        fi
    fi
    return 0
}

function start() {
    check_pid
    running=$?
    if [ $running -gt 0 ];then
        echo -n "$app now is running already, pid="
        cat $pidfile
        return 1
    fi

    if ! [ -f $conf ];then
        echo "Config file $conf doesn't exist, creating one."
        cp cfg.example.json $conf
    fi
    nohup ./$app >/dev/null 2>&1 &
    sleep 1
    running=`ps -p $! | grep -v "PID TTY" | wc -l`
    if [ $running -gt 0 ];then
        echo $! > $pidfile
        echo "$app started..., pid=$!"
    else
        echo "$app failed to start."
        return 1
    fi
}

function stop() {
    pid=`cat $pidfile`
    kill $pid
    rm -f $pidfile
    echo "$app stoped..."
}

function restart() {
    stop
    sleep 1
    start
}

function status() {
    check_pid
    running=$?
    if [ $running -gt 0 ];then
        echo started
    else
        echo stoped
    fi
}

function tailf() {
    tail -f $logfile
}

function build() {
    go build -o $app
    sc=$?
    if [ $sc -ne 0 ];then
        echo "build error"
        exit $sc
    else
        echo "build ok"
    fi
}

function update_gitversion() {
    git log -1 --pretty=%h > $gitversion
}

function pack() {
    build
    version=`./$app -v`
    tar zcvf $app-$version.tar.gz control.sh $app cfg.json.bak $gitversion
}

function packbin() {
    build
    version=`./$app -v`
    tar zcvf $app-bin-$version.tar.gz $app $gitversion
}

function help() {
    echo "$0 build|pack|start|stop|restart|status|tail"
}

if [ "$1" == "" ]; then
    help
elif [ "$1" == "stop" ];then
    stop
elif [ "$1" == "start" ];then
    start
elif [ "$1" == "restart" ];then
    restart
elif [ "$1" == "status" ];then
    status
elif [ "$1" == "tail" ];then
    tailf
elif [ "$1" == "build" ];then
    build
elif [ "$1" == "pack" ];then
    pack
elif [ "$1" == "packbin" ];then
    packbin
else
    help
fi

