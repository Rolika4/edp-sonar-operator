# sonar-operator

![Version: 2.11.0-SNAPSHOT](https://img.shields.io/badge/Version-2.11.0--SNAPSHOT-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.11.0-SNAPSHOT](https://img.shields.io/badge/AppVersion-2.11.0--SNAPSHOT-informational?style=flat-square)

A Helm chart for EDP Sonar Operator

**Homepage:** <https://epam.github.io/edp-install/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| epmd-edp | SupportEPMD-EDP@epam.com | https://solutionshub.epam.com/solution/epam-delivery-platform |
| sergk |  | https://github.com/SergK |

## Source Code

* <https://github.com/epam/edp-sonar-operator>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| annotations | object | `{}` |  |
| global.dnsWildCard | string | `"example.com"` |  |
| global.edpName | string | `""` |  |
| global.openshift.deploymentType | string | `"deploymentConfigs"` |  |
| global.platform | string | `"openshift"` |  |
| image.name | string | `"epamedp/sonar-operator"` |  |
| image.version | string | `nil` |  |
| imagePullPolicy | string | `"IfNotPresent"` |  |
| name | string | `"sonar-operator"` |  |
| nodeSelector | object | `{}` |  |
| resources.limits.memory | string | `"192Mi"` |  |
| resources.requests.cpu | string | `"50m"` |  |
| resources.requests.memory | string | `"64Mi"` |  |
| sonar.affinity | object | `{}` |  |
| sonar.annotations | object | `{}` |  |
| sonar.basePath | string | `""` |  |
| sonar.db.affinity | object | `{}` |  |
| sonar.db.annotations | object | `{}` |  |
| sonar.db.image | string | `"postgres:9.6"` |  |
| sonar.db.imagePullPolicy | string | `"IfNotPresent"` |  |
| sonar.db.nodeSelector | object | `{}` |  |
| sonar.db.resources.limits.memory | string | `"512Mi"` |  |
| sonar.db.resources.requests.cpu | string | `"50m"` |  |
| sonar.db.resources.requests.memory | string | `"64Mi"` |  |
| sonar.db.tolerations | list | `[]` |  |
| sonar.deploy | bool | `true` |  |
| sonar.image | string | `"sonarqube"` |  |
| sonar.imagePullPolicy | string | `"IfNotPresent"` |  |
| sonar.imagePullSecrets | string | `nil` |  |
| sonar.ingress.annotations | object | `{}` |  |
| sonar.ingress.pathType | string | `"Prefix"` |  |
| sonar.initContainers.resources | object | `{}` |  |
| sonar.initImage | string | `"busybox:1.35.0"` |  |
| sonar.name | string | `"sonar"` |  |
| sonar.nodeSelector | object | `{}` |  |
| sonar.plugins.install[0] | string | `"https://github.com/vaulttec/sonar-auth-oidc/releases/download/v2.0.0/sonar-auth-oidc-plugin-2.0.0.jar"` |  |
| sonar.plugins.install[1] | string | `"https://github.com/checkstyle/sonar-checkstyle/releases/download/9.0.1/checkstyle-sonar-plugin-9.0.1.jar"` |  |
| sonar.plugins.install[2] | string | `"https://github.com/spotbugs/sonar-findbugs/releases/download/4.0.4/sonar-findbugs-plugin-4.0.4.jar"` |  |
| sonar.plugins.install[3] | string | `"https://github.com/jensgerdes/sonar-pmd/releases/download/3.3.1/sonar-pmd-plugin-3.3.1.jar"` |  |
| sonar.plugins.install[4] | string | `"https://github.com/sbaudoin/sonar-ansible/releases/download/v2.4.0/sonar-ansible-plugin-2.4.0.jar"` |  |
| sonar.plugins.install[5] | string | `"https://github.com/sbaudoin/sonar-yaml/releases/download/v1.6.0/sonar-yaml-plugin-1.6.0.jar"` |  |
| sonar.plugins.install[6] | string | `"https://github.com/Inform-Software/sonar-groovy/releases/download/1.8/sonar-groovy-plugin-1.8.jar"` |  |
| sonar.resources.limits.memory | string | `"3Gi"` |  |
| sonar.resources.requests.cpu | string | `"100m"` |  |
| sonar.resources.requests.memory | string | `"1.5Gi"` |  |
| sonar.sonarqubeFolder | string | `"/opt/sonarqube"` |  |
| sonar.storage.data.class | string | `"gp2"` |  |
| sonar.storage.data.size | string | `"1Gi"` |  |
| sonar.storage.database.class | string | `"gp2"` |  |
| sonar.storage.database.size | string | `"1Gi"` |  |
| sonar.tolerations | list | `[]` |  |
| sonar.version | string | `"8.9.7-community"` |  |
| tolerations | list | `[]` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.6.0](https://github.com/norwoodj/helm-docs/releases/v1.6.0)