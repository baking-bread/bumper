package version

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type SemVer struct {
	Major  int
	Minor  int
	Patch  int
	log    *logrus.Logger
	dryRun bool
	force  bool
}

func NewSemVer(log *logrus.Logger, dryRun bool, force bool) *SemVer {
	return &SemVer{
		log:    log,
		dryRun: dryRun,
		force:  force,
	}
}

func (sv *SemVer) Parse(version string) error {
	version = strings.TrimPrefix(version, "v")
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return errors.New("invalid version format")
	}

	var err error
	if sv.Major, err = strconv.Atoi(parts[0]); err != nil {
		return err
	}
	if sv.Minor, err = strconv.Atoi(parts[1]); err != nil {
		return err
	}
	if sv.Patch, err = strconv.Atoi(parts[2]); err != nil {
		return err
	}

	return nil
}

func (sv *SemVer) Compare(other *SemVer) int {
	if sv.Major != other.Major {
		return sv.Major - other.Major
	}
	if sv.Minor != other.Minor {
		return sv.Minor - other.Minor
	}
	return sv.Patch - other.Patch
}

func (sv *SemVer) Bump(bumpType string) error {
	oldVer := &SemVer{
		Major: sv.Major,
		Minor: sv.Minor,
		Patch: sv.Patch,
	}

	newVer := &SemVer{
		Major: sv.Major,
		Minor: sv.Minor,
		Patch: sv.Patch,
	}

	switch bumpType {
	case "patch":
		newVer.Patch++
	case "minor":
		newVer.Minor++
		newVer.Patch = 0
	case "major":
		newVer.Major++
		newVer.Minor = 0
		newVer.Patch = 0
	default:
		return errors.New("invalid bump type")
	}

	if newVer.Compare(oldVer) < 0 && !sv.force {
		sv.log.Errorf("Attempting to downgrade version from %s to %s. Use --force to override.", oldVer.String(), newVer.String())
		return errors.New("version downgrade detected. Use --force to override")
	}

	if sv.dryRun {
		sv.log.Info("Dry-run mode enabled. No changes will be applied.")
		sv.log.Infof("Would bump version to: %s", newVer.String())
		return nil
	}

	sv.Major = newVer.Major
	sv.Minor = newVer.Minor
	sv.Patch = newVer.Patch
	sv.log.Infof("Bumped version to: %s", sv.String())
	return nil
}

func (sv *SemVer) String() string {
	return fmt.Sprintf("v%d.%d.%d", sv.Major, sv.Minor, sv.Patch)
}
