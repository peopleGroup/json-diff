const fs = require('fs');
const goLivePath = './log/live'
const goLocalPath = './log/local'
const diffFilePath = './log/diff'
const decorators = fs.readdirSync(goLivePath)

const dataProvider = (decorator, fileName) => {
    const localResponsePath = `${goLocalPath}/${decorator}/${fileName}`
    const liveResponsePath = `${goLivePath}/${decorator}/${fileName}`
    const localData = parse(read(localResponsePath))
    const liveData = parse(read(liveResponsePath))
    return [localData, liveData]
}

const adm = () => ({ added: [], modified: [], deleted: [] })
const isPrimitive = v => v !== Object(v)
const intersection = (xs, ys) => xs.filter(x => ys.includes(x))
const setDifference = (xs, ys) => {
    const common = intersection(xs, ys)
    return xs.filter(x => !common.includes(x))
}
const parse = JSON.parse
const keys = Object.keys
const read = path => fs.readFileSync(path, 'utf8')
const uniq = xs => keys(xs.reduce((acc, x) => acc[x] = true && acc, {}))
const isEmpty = xs => xs.length === 0

const diff = (o1, o2, path, diffAdm) => {
    if (!o1 || !o2) return diffAdm
    const [keys1, keys2] = [keys(o1), keys(o2)]
    const [addedKeys, deletedKeys] = [setDifference(keys2, keys1), setDifference(keys1, keys2)]
    diffAdm.added = [...diffAdm.added, ...addedKeys.map(k => `${path}.${k}`)]
    diffAdm.deleted = [...diffAdm.deleted, ...deletedKeys.map(k => `${path}.${k}`)]
    const modifiedKeys = intersection(keys1, keys2)
    diffAdm = modifiedKeys.reduce((acc, key) => {
        const fullKey = `${path}.${key}`
        const [v1, v2] = [o1[key], o2[key]]
        if (isPrimitive(v1) && isPrimitive(v2)) {
            if (v1 !== v2) {
                const diffKey = fullKey.split('.').slice(3).join('.')
                acc.modified = [...acc.modified, diffKey]
            }
            return acc
        } else {
            return diff(v1, v2, fullKey, diffAdm)
        }
    }, diffAdm)
    return diffAdm
}

const writeChangesToFile = (fileName, added, modified, deleted) => {
    const data = `
    Added Keys:     \n${added.reduce((acc, x) => `${acc}\n\t\t${x}`, '')}\n
    Modified Keys:  \n${modified.reduce((acc, x) => `${acc}\n\t\t${x}`, '')}\n
    Deleted Keys:   \n${deleted.reduce((acc, x) => `${acc}\n\t\t${x}`, '')}\n
    `;
    fs.writeFileSync(`${diffFilePath}/${fileName}.txt`, data)
};

decorators.forEach(decorator => {
    const files = fs.readdirSync(`${goLivePath}/${decorator}/`)
    files.forEach(file => {
        const [localData, liveData] = dataProvider(decorator, file)
        let changes = diff(localData, liveData, '', adm())
        let { added, modified, deleted } = changes

        if (!isEmpty(added) || !isEmpty(modified) || !isEmpty(deleted)) {
            writeChangesToFile(`${decorator}/${file}`, added, modified, deleted)
        }
    })
});