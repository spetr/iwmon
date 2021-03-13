# IWmon

IWMon is IceWarp server external monitoring tool with Prometheus exporter API and Zabbix sender. All configuration
is stored in "iwmon.yml" configuration file located in same directory with the main binary.

Install on Linux:

```yaml
mkdir /opt/iwmon
tar -C /opt/iwmon -xzf iwmon_v0.2.6_linux_amd64.tgz
cd /opt/iwmon
./iwmon -service install
./iwmon -service start
```

Install on Windows:

Unzip to C:\iwmon and run cmd.exe aith administrator account.

```cmd
cd c:\iwmon
iwmon.exe -service install
iwmon.exe -service start
```
