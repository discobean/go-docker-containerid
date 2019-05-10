package containerid

import (
	"errors"
	"io/ioutil"
	"regexp"
)

func Find() (string, error) {
	data, err := ioutil.ReadFile("/proc/self/cgroup")
	if err != nil {
		return "", err
	}

	// search for docker/[a-z0-9]{12}
	re := regexp.MustCompile("docker/([a-z0-9]{12})")
	results := re.FindSubmatch([]byte(data))
	if results == nil {
		return "", errors.New("no matching docker process in cgroup found")
	}

	return string(results[1]), nil
}