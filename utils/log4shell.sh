#!/bin/bash

source ./b-log.sh

function log4shell(){
    # logging shell
    # 1. resave current log to an old log file
    # 2. remove too old log file


    str_head=`cat $log_path|head -1`
    echo $str_head

    crtime=`echo "$str_head" | awk  -F'[' '{print $3}'|tr -d '[]'`
    # echo $crtime
    crday=`echo $crtime| awk '{print $1}'`
    # echo $crday

    echo 'log_path --> '$log_path

    # save log 
    if [ $log_save_day -gt 0 ]; then
        presaveday=`date -d "$log_save_day days ago" +%Y-%m-%d`
        # echo 'presaveday --> '$presaveday

        if [ "$crday" != "" ] && [ "$presaveday" == "$crday" ]; then
            mv $log_path $log_path.$presaveday
            INFO "mv $log_path $log_path.$presaveday" >> $log_path
        fi
    fi

    # clear log
    if [ $log_clear_day -gt 0 ]; then
        preclearday=`date -d "$log_clear_day days ago" +%Y-%m-%d`
        # echo 'preclearday --> '$preclearday

        if [ -f "$log_path.$preclearday" ]; then
            rm $log_path.$preclearday
            INFO "rm $log_path.$preclearday" >> $log_path
        fi
    fi
}