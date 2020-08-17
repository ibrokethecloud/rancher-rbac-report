## kubectl-rancher-rbac-report 

Simple plugin to generate rbac report for your Rancher installation.

Use is as follows:

```
The plugin interacts with the k8s api of the cluster where Rancher is installed (local) cluster, and attempts to
generate a list of all rbac settings being managed by Rancher.

Usage:
  kubectl-rancher-rbac-report [flags]

Flags:
  -c, --cluster string   Generate report for specific cluster only
  -h, --help             help for kubectl-rancher-rbac-report
```

A sample report would look as follows.

```
+--------------------------------+-------------+
|              USER              | GLOBAL ROLE |
+--------------------------------+-------------+
| Default Admin                  | admin       |
| System account for Cluster     | user        |
| local                          |             |
| System account for Project     | user        |
| p-m2zpj                        |             |
| System account for Project     | user        |
| p-zpwl8                        |             |
|                                | user        |
| System account for Cluster     | user        |
| c-7l8tl                        |             |
| System account for Project     | user        |
| p-flqgs                        |             |
| System account for Project     | user        |
| p-t6t85                        |             |
| System account for Project     | user        |
| p-6wpc8                        |             |
| System account for Project     | user        |
| p-m49kl                        |             |
| System account for Project     | user        |
| p-t9tlr                        |             |
| System account for Project     | user        |
| p-bpr4v                        |             |
|                                | user        |
| System account for Project     | user        |
| p-fl8t4                        |             |
| System account for Project     | user        |
| p-wjh6s                        |             |
+--------------------------------+-------------+
+---------+--------------------------------+-------+--------------------+---------------+
| CLUSTER |              USER              | GROUP | GROUP PRINCIPAL ID | CLUSTER ROLE  |
+---------+--------------------------------+-------+--------------------+---------------+
| local   | Default Admin                  |       |                    | cluster-owner |
| local   | System account for Cluster     |       |                    | cluster-owner |
|         | local                          |       |                    |               |
+---------+--------------------------------+-------+--------------------+---------------+
+---------+---------+--------------------------------+-------+--------------------+----------------+
| CLUSTER | PROJECT |              USER              | GROUP | GROUP PRINCIPAL ID |  PROJECT ROLE  |
+---------+---------+--------------------------------+-------+--------------------+----------------+
| local   | System  | Default Admin                  |       |                    | project-owner  |
| local   | System  | System account for Project     |       |                    | project-member |
|         |         | p-zpwl8                        |       |                    |                |
| local   | Default | System account for Project     |       |                    | project-member |
|         |         | p-m2zpj                        |       |                    |                |
| local   | Default | Default Admin                  |       |                    | project-owner  |
+---------+---------+--------------------------------+-------+--------------------+----------------+
+---------+--------------------------------+-------+--------------------+---------------+
| CLUSTER |              USER              | GROUP | GROUP PRINCIPAL ID | CLUSTER ROLE  |
+---------+--------------------------------+-------+--------------------+---------------+
| vmware  | System account for Cluster     |       |                    | cluster-owner |
|         | c-7l8tl                        |       |                    |               |
| vmware  | Default Admin                  |       |                    | cluster-owner |
| vmware  |                                |       |                    | cluster-owner |
+---------+--------------------------------+-------+--------------------+---------------+
+---------+-----------+--------------------------------+-------+--------------------+----------------+
| CLUSTER |  PROJECT  |              USER              | GROUP | GROUP PRINCIPAL ID |  PROJECT ROLE  |
+---------+-----------+--------------------------------+-------+--------------------+----------------+
| vmware  | System    | System account for Project     |       |                    | project-member |
|         |           | p-flqgs                        |       |                    |                |
| vmware  | System    | Default Admin                  |       |                    | project-owner  |
| vmware  | Default   | Default Admin                  |       |                    | project-owner  |
| vmware  | Default   | System account for Project     |       |                    | project-member |
|         |           | p-t6t85                        |       |                    |                |
| vmware  | kubeless  | Default Admin                  |       |                    | project-owner  |
| vmware  | kubeless  | System account for Project     |       |                    | project-member |
|         |           | p-6wpc8                        |       |                    |                |
| vmware  | enforcer  | Default Admin                  |       |                    | project-owner  |
| vmware  | enforcer  | System account for Project     |       |                    | project-member |
|         |           | p-m49kl                        |       |                    |                |
| vmware  | tools     | Default Admin                  |       |                    | project-owner  |
| vmware  | tools     | System account for Project     |       |                    | project-member |
|         |           | p-t9tlr                        |       |                    |                |
| vmware  | hobbyfarm | Default Admin                  |       |                    | project-owner  |
| vmware  | hobbyfarm | System account for Project     |       |                    | project-member |
|         |           | p-bpr4v                        |       |                    |                |
| vmware  | data      | Default Admin                  |       |                    | project-owner  |
| vmware  | data      | System account for Project     |       |                    | project-member |
|         |           | p-fl8t4                        |       |                    |                |
| vmware  | data      |                                |       |                    | CustomTest     |
| vmware  | data      | Default Admin                  |       |                    | CustomTest     |
| vmware  | bookinfo  | Default Admin                  |       |                    | project-owner  |
| vmware  | bookinfo  | System account for Project     |       |                    | project-member |
|         |           | p-wjh6s                        |       |                    |                |
+---------+-----------+--------------------------------+-------+--------------------+----------------+
```