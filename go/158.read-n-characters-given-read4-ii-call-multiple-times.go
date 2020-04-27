/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
    var localBuffer []byte
    var localBufferIndex int
    var localBufferLength int
    var bufferIndex int // deprecated: turns out that we dont need to memorize the buffer cursor
    var eof bool
    return func(buf []byte, n int) int {
        if eof {
            if localBuffer == nil {
                return 0
            } else {
                count := 0
                for count < n && localBufferIndex < localBufferLength {
                    // buf[bufferIndex] = localBuffer[localBufferIndex]
                    buf[count] = localBuffer[localBufferIndex]
                    bufferIndex++
                    localBufferIndex++
                    count++
                }
                if localBufferIndex == localBufferLength {
                    localBuffer = nil
                    localBufferIndex = 0
                    localBufferLength = 0
                }
                return count
            }
        } else {
            count := 0
            // read from local buffer first
            if localBuffer != nil {
                for localBufferIndex < localBufferLength && count < n {
                    // buf[bufferIndex] = localBuffer[localBufferIndex]
                    buf[count] = localBuffer[localBufferIndex]
                    bufferIndex++
                    localBufferIndex++
                    count++
                }
            }
            
            if count == n {
                return count
            }
           
            // calculate required time
            readTime := (n-count)/4
            if (n-count) % 4 != 0 {
                readTime++
            }
           
            // read all complete blocks (4 chars)
            for i := 0; i < readTime - 1; i++ {
                tmp := make([]byte, 4)
                readCount := read4(tmp)
                for j := 0; j < readCount; j++ {
                    // buf[bufferIndex] = tmp[j]
                    buf[count] = tmp[j]
                    bufferIndex++
                    count++
                }
                if readCount < 4 {
                    // end of file
                    eof = true
                    localBuffer = nil
                    localBufferIndex = 0
                    localBufferLength = 0
                    return count
                }
            }
            
            // read the last required blocks
            localBuffer = make([]byte, 4)
            localBufferIndex = 0
            localBufferLength = read4(localBuffer)
            if localBufferLength < 4 {
                eof = true
            }
            for localBufferIndex < localBufferLength && count < n {
                // buf[bufferIndex] = localBuffer[localBufferIndex]
                buf[count] = localBuffer[localBufferIndex]
                bufferIndex++
                localBufferIndex++
                count++
            }
            return count
        }
    }
}