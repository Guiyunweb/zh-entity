export default ({request}) => ({
    CONVERSION_JAVA(data = {}) {
        return request({
            url: '/conversion/java',
            method: 'post',
            data
        })
    },
    CONVERSION_SQL(data = {}) {
        return request({
            url: '/conversion/sql',
            method: 'post',
            data
        })
    }
})
