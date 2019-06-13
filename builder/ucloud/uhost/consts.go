package uhost

const (
	// defaultPasswordStr, defaultPasswordNum and defaultPasswordSpe are used to general default value of root password of UHost instance
	defaultPasswordNum = "012346789"
	defaultPasswordStr = "abcdefghijklmnopqrstuvwxyz"
	defaultPasswordSpe = "-_"
)

var bootDiskTypeMap = map[string]string{
	"cloud_ssd":    "CLOUD_SSD",
	"local_normal": "LOCAL_NORMAL",
	"local_ssd":    "LOCAL_SSD",
}
