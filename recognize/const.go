package recognize

const UriComponentVar = "{component}"
const UriTypeVar = "{type}"
const UriVersionVar = "{version}"

const WpComponentDownloadUriTemplate = "https://downloads.wordpress.org/{type}/{component}.{version}.zip"
const WpComponentRepoUrlRegex = "^https?://wordpress.org/{type}s/([^/]+)/?"
