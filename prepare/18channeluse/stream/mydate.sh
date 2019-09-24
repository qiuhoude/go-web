#!/bin/bash

params=(${@})

    if [[ ${#params[*]} -gt 0 ]] ;then
            for i in ${params[@]} ;do
                if [[ "$i" = "-s" ]] ;then
                    while [[ "${confirm}" != "yes" ]] && [[ "${confirm}" != "no" ]]; do
                        echo -ne "\e[4;41;37m 线上修改时间谨慎操作!!! \e[0m\n 是否继续 (yes or no):"
                        read confirm
                    done
                    [[ "${confirm}" == "yes" ]] && /bin/date $@
                    exit
                fi
            done
        fi


# 直接执行
/bin/date $@

