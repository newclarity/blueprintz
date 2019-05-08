package global

type BlueprintType = string
type PartName = string

type Url = string
type Slug = string
type Version = string
type RelativeDir = string
type Domain = string

type AbsoluteDirs []AbsoluteDir
type AbsoluteDir = string


type Layouter interface {
	GetProjectDir() RelativeDir
	GetWebrootDir() RelativeDir
	GetContentDir() RelativeDir
	GetCoreDir() RelativeDir   
}