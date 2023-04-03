package models

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/stackrox/release-registry/pkg/configuration"
	"github.com/stackrox/release-registry/pkg/logging"
	"github.com/stackrox/release-registry/pkg/storage"
	"github.com/stackrox/release-registry/pkg/utils/validate"
)

//nolint:gochecknoglobals
var log = logging.CreateProductionLogger()

const (
	errorCannotCreateRelease = "can not create release"
	errorInvalidVersion      = "invalid version specified"
)

// CreateRelease creates a new Release based on based information.
func CreateRelease(
	config *configuration.Config,
	tag, commit, creator string, metadata []ReleaseMetadata,
) (*Release, error) {
	if err := validate.IsValidVersion(tag); err != nil {
		return nil, errors.Wrap(err, errorCannotCreateRelease)
	}

	if err := validate.IsValidCommit(commit); err != nil {
		return nil, errors.Wrap(err, errorCannotCreateRelease)
	}

	if err := validate.IsValidActorEmail(config, creator); err != nil {
		return nil, errors.Wrap(err, errorCannotCreateRelease)
	}

	release := &Release{
		Tag:      tag,
		Commit:   commit,
		Creator:  creator,
		Metadata: metadata,
	}
	result := storage.DB.Where(release).FirstOrCreate(release)

	if result.Error != nil {
		return nil, result.Error
	}

	log.Infow("release created", "creator", release.Creator, "tag", release.Tag)

	return release, nil
}

// RejectRelease rejects a release identified by tag.
func RejectRelease(tag string, preload bool) (*Release, error) {
	if err := validate.IsValidVersion(tag); err != nil {
		return nil, errors.Wrap(err, errorInvalidVersion)
	}

	release, err := GetRelease(tag, preload, false)
	if err != nil {
		return nil, errors.Wrap(err, "release not found or already rejected")
	}

	result := storage.DB.Model(release).Update("rejected", true)
	if result.Error != nil {
		return nil, result.Error
	}

	log.Infow("release rejected", "tag", release.Tag)

	return release, nil
}

// GetRelease returns a Release to a tag.
func GetRelease(tag string, preload, includeRejected bool) (*Release, error) {
	if err := validate.IsValidVersion(tag); err != nil {
		return nil, errors.Wrap(err, errorInvalidVersion)
	}

	release := &Release{}
	tx := storage.DB.Where("tag = ?", tag)
	tx = withPreloadedMetadata(tx, preload)
	tx = withPreloadedQualityMilestones(tx, preload)
	tx = withPreloadedQualityMilestoneDefinitions(tx, preload)
	tx = withIncludedRejectedReleases(tx, includeRejected)

	result := tx.First(release)
	if result.Error != nil {
		return nil, result.Error
	}

	return release, nil
}

// ListAllReleases returns all known Releases.
func ListAllReleases(preload bool, includeRejected bool) ([]Release, error) {
	releases := []Release{}
	tx := storage.DB
	tx = withPreloadedMetadata(tx, preload)
	tx = withPreloadedQualityMilestones(tx, preload)
	tx = withPreloadedQualityMilestoneDefinitions(tx, preload)
	tx = withIncludedRejectedReleases(tx, includeRejected)

	result := tx.Find(&releases)
	if result.Error != nil {
		return nil, result.Error
	}

	return releases, nil
}

// ListAllReleasesWithPrefix implements search to return all Releases starting with a specific prefix.
func ListAllReleasesWithPrefix(prefix string, preload, includeRejected bool) ([]Release, error) {
	if err := validate.IsNotEmpty(prefix); err != nil {
		return nil, errors.Wrapf(err, "prefix parameter is empty")
	}

	releases := []Release{}
	tx := storage.DB.Where("tag LIKE ?", fmt.Sprintf("%s%%", prefix))
	tx = withPreloadedMetadata(tx, preload)
	tx = withPreloadedQualityMilestones(tx, preload)
	tx = withPreloadedQualityMilestoneDefinitions(tx, preload)
	tx = withIncludedRejectedReleases(tx, includeRejected)

	result := tx.Find(&releases)
	if result.Error != nil {
		return nil, result.Error
	}

	return releases, nil
}

// ListAllReleasesAtQualityMilestone returns all Releases that have reached a specific QualityMilestone.
func ListAllReleasesAtQualityMilestone(qualityMilestoneName string, preload, includeRejected bool) ([]Release, error) {
	if err := validate.IsNotEmpty(qualityMilestoneName); err != nil {
		return nil, errors.Wrapf(err, "qualityMilestoneName parameter is empty")
	}

	releases := []Release{}
	tx := storage.DB.Where("quality_milestone_definitions.name = ?", qualityMilestoneName)
	tx = joinReleasesWithQualityMilestoneDefinitions(tx)
	tx = withPreloadedMetadata(tx, preload)
	tx = withPreloadedQualityMilestones(tx, preload)
	tx = withPreloadedQualityMilestoneDefinitions(tx, preload)
	tx = withIncludedRejectedReleases(tx, includeRejected)

	result := tx.Find(&releases)
	if result.Error != nil {
		return nil, result.Error
	}

	return releases, nil
}

// ListAllReleasesWithPrefixAtQualityMilestone implements search to return all Releases starting
// with a specific prefix at a specific QualityMilestone.
func ListAllReleasesWithPrefixAtQualityMilestone(
	prefix, qualityMilestoneName string,
	preload, includeRejected bool,
) ([]Release, error) {
	if err := validate.IsNotEmpty(prefix); err != nil {
		return nil, errors.Wrapf(err, "prefix parameter is empty")
	}

	if err := validate.IsNotEmpty(qualityMilestoneName); err != nil {
		return nil, errors.Wrapf(err, "qualityMilestoneName parameter is empty")
	}

	releases := []Release{}

	tx := joinReleasesWithQualityMilestoneDefinitions(storage.DB)
	tx = tx.Where("quality_milestone_definitions.name = ?", qualityMilestoneName)
	tx = tx.Where("releases.tag LIKE ?", fmt.Sprintf("%s%%", prefix))
	tx = withPreloadedMetadata(tx, preload)
	tx = withPreloadedQualityMilestones(tx, preload)
	tx = withPreloadedQualityMilestoneDefinitions(tx, preload)
	tx = withIncludedRejectedReleases(tx, includeRejected)

	result := tx.Find(&releases)
	if result.Error != nil {
		return nil, result.Error
	}

	return releases, nil
}

// FindLatestRelease returns the latest Release overall, sorted by semantic versioning.
func FindLatestRelease(preload, includeRejected bool) (*Release, error) {
	releases, err := ListAllReleases(false, includeRejected)
	if err != nil {
		return nil, err
	}

	latestVersion, err := findLatestVersionFromListOfReleases(releases)
	if err != nil {
		return nil, err
	}

	return GetRelease(latestVersion, preload, includeRejected)
}

// FindLatestReleaseWithPrefix returns the latest Release with a prefix, sorted by semantic versioning.
func FindLatestReleaseWithPrefix(prefix string, preload, includeRejected bool) (*Release, error) {
	releases, err := ListAllReleasesWithPrefix(prefix, false, includeRejected)
	if err != nil {
		return nil, err
	}

	latestVersion, err := findLatestVersionFromListOfReleases(releases)
	if err != nil {
		return nil, err
	}

	return GetRelease(latestVersion, preload, includeRejected)
}

// FindLatestReleaseAtQualityMilestone returns the latest Release at a QualityMilestone, sorted by semantic versioning.
func FindLatestReleaseAtQualityMilestone(qualityMilestoneName string, preload, includeRejected bool) (*Release, error) {
	releases, err := ListAllReleasesAtQualityMilestone(qualityMilestoneName, false, includeRejected)
	if err != nil {
		return nil, err
	}

	latestVersion, err := findLatestVersionFromListOfReleases(releases)
	if err != nil {
		return nil, err
	}

	return GetRelease(latestVersion, preload, includeRejected)
}

// FindLatestRelaseWithPrefixAtQualityMilestone returns the latest Release with a prefix at a QualityMilestone,
// sorted by semantic versioning.
func FindLatestRelaseWithPrefixAtQualityMilestone(
	prefix, qualityMilestoneName string,
	preload, includeRejected bool,
) (*Release, error) {
	releases, err := ListAllReleasesWithPrefixAtQualityMilestone(prefix, qualityMilestoneName, false, includeRejected)
	if err != nil {
		return nil, err
	}

	latestVersion, err := findLatestVersionFromListOfReleases(releases)
	if err != nil {
		return nil, err
	}

	return GetRelease(latestVersion, preload, includeRejected)
}