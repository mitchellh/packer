module github.com/hashicorp/packer

require (
	cloud.google.com/go v0.66.0
	github.com/1and1/oneandone-cloudserver-sdk-go v1.0.1
	github.com/Azure/azure-sdk-for-go v40.5.0+incompatible
	github.com/Azure/go-autorest/autorest v0.10.0
	github.com/Azure/go-autorest/autorest/adal v0.8.2
	github.com/Azure/go-autorest/autorest/azure/auth v0.4.2
	github.com/Azure/go-autorest/autorest/azure/cli v0.3.1
	github.com/Azure/go-autorest/autorest/date v0.2.0
	github.com/Azure/go-autorest/autorest/to v0.3.0
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/ChrisTrenkamp/goxpath v0.0.0-20170922090931-c385f95c6022
	github.com/NaverCloudPlatform/ncloud-sdk-go-v2 v1.1.0
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/Telmate/proxmox-api-go v0.0.0-20200715182505-ec97c70ba887
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190418113227-25233c783f4e
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20170113022742-e6dbea820a9f
	github.com/antihax/optional v1.0.0
	github.com/approvals/go-approval-tests v0.0.0-20160714161514-ad96e53bea43
	github.com/aws/aws-sdk-go v1.37.15
	github.com/biogo/hts v0.0.0-20160420073057-50da7d4131a3
	github.com/c2h5oh/datasize v0.0.0-20200112174442-28bbd4740fee
	github.com/cheggaaa/pb v1.0.27
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/digitalocean/go-qemu v0.0.0-20201211181942-d361e7b4965f
	github.com/digitalocean/godo v1.11.1
	github.com/dylanmei/iso8601 v0.1.0 // indirect
	github.com/exoscale/packer-plugin-exoscale v0.1.0
	github.com/fatih/camelcase v1.0.0
	github.com/fatih/structtag v1.0.0
	github.com/go-ini/ini v1.25.4
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-resty/resty/v2 v2.3.0
	github.com/gobwas/glob v0.2.3
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3
	github.com/google/go-cmp v0.5.2
	github.com/google/go-github/v33 v33.0.1-0.20210113204525-9318e629ec69
	github.com/google/uuid v1.1.2
	github.com/gophercloud/gophercloud v0.12.0
	github.com/gophercloud/utils v0.0.0-20200508015959-b0167b94122c
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/hako/durafmt v0.0.0-20200710122514-c0fb7b4da026
	github.com/hashicorp/aws-sdk-go-base v0.6.0
	github.com/hashicorp/errwrap v1.0.0
	github.com/hashicorp/go-checkpoint v0.0.0-20171009173528-1545e56e46de
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/go-cty-funcs v0.0.0-20200930094925-2721b1e36840
	github.com/hashicorp/go-getter/v2 v2.0.0-20200604122502-a6995fa1edad
	github.com/hashicorp/go-multierror v1.1.0
	github.com/hashicorp/go-oracle-terraform v0.0.0-20181016190316-007121241b79
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/go-version v1.2.0
	github.com/hashicorp/hcl/v2 v2.8.0
	github.com/hashicorp/packer-plugin-docker v0.0.2
	github.com/hashicorp/packer-plugin-sdk v0.0.14
	github.com/hashicorp/vault/api v1.0.4
	github.com/hetznercloud/hcloud-go v1.15.1
	github.com/hyperonecom/h1-client-go v0.0.0-20191203060043-b46280e4c4a4
	github.com/jdcloud-api/jdcloud-sdk-go v1.9.1-0.20190605102154-3d81a50ca961
	github.com/joyent/triton-go v0.0.0-20180628001255-830d2b111e62
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/klauspost/compress v1.11.12 // indirect
	github.com/klauspost/crc32 v1.2.0 // indirect
	github.com/klauspost/pgzip v0.0.0-20151221113845-47f36e165cec
	github.com/kr/fs v0.1.0 // indirect
	github.com/linode/linodego v0.14.0
	github.com/masterzen/winrm v0.0.0-20201030141608-56ca5c5f2380
	github.com/mattn/go-tty v0.0.0-20191112051231-74040eebce08
	github.com/mitchellh/cli v1.1.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-vnc v0.0.0-20150629162542-723ed9867aed
	github.com/mitchellh/gox v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.4.0
	github.com/mitchellh/panicwrap v1.0.0
	github.com/mitchellh/prefixedio v0.0.0-20151214002211-6e6954073784
	github.com/mitchellh/reflectwalk v1.0.0
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/olekukonko/tablewriter v0.0.0-20180105111133-96aac992fc8b
	github.com/oracle/oci-go-sdk v24.3.0+incompatible
	github.com/outscale/osc-sdk-go/osc v0.0.0-20200722135656-d654809d0699
	github.com/pierrec/lz4 v2.0.5+incompatible
	github.com/pkg/errors v0.9.1
	github.com/posener/complete v1.2.3
	github.com/profitbricks/profitbricks-sdk-go v4.0.2+incompatible
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.7
	github.com/shirou/gopsutil v3.21.1+incompatible
	github.com/stretchr/testify v1.7.0
	github.com/tencentcloud/tencentcloud-sdk-go v3.0.222+incompatible
	github.com/ucloud/ucloud-sdk-go v0.16.3
	github.com/ufilesdk-dev/ufile-gosdk v0.0.0-20190830075812-b4dbc4ef43a6
	github.com/ulikunitz/xz v0.5.5
	github.com/vmware/govmomi v0.23.1
	github.com/xanzy/go-cloudstack v0.0.0-20190526095453-42f262b63ed0
	github.com/yandex-cloud/go-genproto v0.0.0-20200915125933-33de72a328bd
	github.com/yandex-cloud/go-sdk v0.0.0-20200921111412-ef15ded2014c
	github.com/zclconf/go-cty v1.7.0
	github.com/zclconf/go-cty-yaml v1.0.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/mobile v0.0.0-20201208152944-da85bec010a2
	golang.org/x/mod v0.3.0
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c
	golang.org/x/tools v0.0.0-20201111133315-69daaf961d65
	google.golang.org/api v0.32.0
	google.golang.org/grpc v1.32.0
	gopkg.in/ini.v1 v1.62.0 // indirect
	inet.af/netaddr v0.0.0-20210317195617-2d42ec05f8a1
)

replace github.com/hashicorp/packer-plugin-sdk => github.com/paginabianca/packer-plugin-sdk v0.1.1

go 1.16
