#!/bin/bash

portname=$1
nsname=$portname
mac=$2

ip netns del $nsname 1>/dev/null 2>&1
ip netns add $nsname
 
veth=veth${portname: -8}
tap=$portname

ip link add $veth address $mac type veth peer name $tap
 
ip link set $veth netns $nsname
 
ip netns exec $nsname ifconfig lo up
ip netns exec $nsname ifconfig $veth up
ip netns exec $nsname ip link set dev $veth
ifconfig $tap up
ovs-vsctl add-port br0 $tap -- set Interface $tap external-ids:iface-status=active -- set Interface $tap external-ids:attached-mac=$mac
 
#ip netns exec $nsname dhclient -r $veth
#ip netns exec $nsname dhclient $veth
#ip netns exec $nsname ifconfig
