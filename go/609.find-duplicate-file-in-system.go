/*
 * @lc app=leetcode id=609 lang=golang
 *
 * [609] Find Duplicate File in System
 */

// @lc code=start
func findDuplicate(paths []string) [][]string {
    contents := make(map[string][]string)
    for _, path := range paths {
        files := strings.Split(path, " ")
        directory := files[0]
        files = files[1:]
        for _, file := range files {
            position := strings.Index(file, ".txt(")
            name := file[:position+4]
            content := file[position+5:len(file)-1]
            contents[content] = append(contents[content], directory + "/" + name)
        }
    }
   
    var result [][]string
    for _, path := range contents {
        if len(path) > 1 {
            result = append(result, path)
        }
    }
    
    return result
}
// @lc code=end

