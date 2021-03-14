# IWmon

IWMon is IceWarp server external monitoring tool with Prometheus exporter API and Zabbix sender. All configuration
is stored in "iwmon.yml" configuration file located in same directory with the main binary.

### Install on Linux:

```bash
mkdir /opt/iwmon
tar -C /opt/iwmon -xzf iwmon_v0.2.7_linux_amd64.tgz
cd /opt/iwmon
./iwmon -service install
./iwmon -service start
```

### Install on Windows:

Unzip to C:\iwmon and run cmd.exe aith administrator account.

```cmd
cd c:\iwmon
iwmon.exe -service install
iwmon.exe -service start
```

### Configuration file example:

```yaml
icewarp:

  # Monitoring using tool.sh
  tool:
    path: "/opt/icewarp/tool.sh"
    timeout: 5                      # Command timeout
    concurrency: 2                  # Max command concurrency

  # Monitoring using snmp
  snmp: 
    address: "0.0.0.0:161"
    timeout: 5

  # Automatic value refresh (in seconds)
  refresh:                          
    version: 3600
    fs_mail: 60
    snmp: 60

zabbix-sender:
  enabled: true
  hostname: icewarp-server.example.com
  servers:
    - 172.16.254.1:10051
```

### Links

https://github.com/milamber86/scripts/blob/master/iwmon_agent.sh
