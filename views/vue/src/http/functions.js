export const buildParamsUri = () => {
    const url = new URL(window.location.href)
    let queryString = url.search.substring(1)
    let buildParams;

    if (queryString) {
        if (queryString.indexOf('&') != -1) {
            const queryParts = queryString.split('&')
            buildParams = {}

            for (let key in queryParts) {
                let item = queryParts[key]
                
                item = item.split('=')

                if (item[0].length) buildParams[item[0]] = item[1]
            }
        } else {
            const queryPart = queryString.split('=')

            if (queryPart[0].length) {
                buildParams = {}
                buildParams[queryPart[0]] = queryPart[1]
            }
        }
    }

    return (typeof buildParams === 'object') ? buildParams : false
}