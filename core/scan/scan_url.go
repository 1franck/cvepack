package scan

//type ScanUrl struct {
//    Url      string
//    Projects []ecosystem.Project
//    Verbose  bool
//}
//
//func NewScanUrl(url string, verbose bool) *ScanUrl {
//    return &ScanUrl{Url: url, Verbose: verbose}
//}
//
//func (scan *ScanUrl) Log(msg string) {
//    if scan.Verbose {
//        fmt.Println(msg)
//    }
//}
//
//func (scan *ScanUrl) Run() []ecosystem.Project {
//    projects := make([]ecosystem.Project, 0)
//    if github.DetectGithubRepoUrl(scan.Url) {
//        scan.Log("Detected GitHub repo")
//        return projects
//    }
//    return projects
//}
