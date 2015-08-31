# Linux

# 目录

## 1. History 增强
        # vi /etc/profile;source /etc/profile
        USER_IP=`who -u am i 2>/dev/null| awk '{print $NF}'|sed -e 's/[()]//g'`
        HISTDIR=/var/log/.hist
        if [ -z $USER_IP ];then
            USER_IP=`hostname`
        fi
        if [ ! -d $HISTDIR ];then
            mkdir -p $HISTDIR
            chmod 777 $HISTDIR
        fi
        if [ ! -d $HISTDIR/${USER_IP} ];then
            mkdir -p $HISTDIR/${USER_IP}
            chmod 300 $HISTDIR/${USER_IP}
        fi
        chmod 600 $HISTDIR/${USER_IP}/*.hist* 2>/dev/null

        DT=`date +%Y-%m-%d_%H:%M:%S`
        HDT=`date +%Y-%m`
        echo "$DT IP:$USER_IP USER:$LOGNAME login" >> $HISTDIR/${USER_IP}/${LOGNAME}.hist.$HDT
        export HISTSIZE=4096
        export HISTFILE="$HISTDIR/${USER_IP}/${LOGNAME}.hist.$HDT"
        export PROMPT_COMMAND="history -a"
        export HISTTIMEFORMAT="[%F %T][`whoami`][${USER_IP}] "