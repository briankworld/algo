/**
 * @param {number[][]} points
 * @return {number}
 */
var minAreaRect = function(points) {
    let s = new Set(points.map(p => `${p[0]},${p[1]}`))
    let min = Infinity

    for (let i =0; i < points.length; i++) {
        for (let j = i + 1; j < points.length; j++) {
            let [x1, y1] = points[i]
            let [x2, y2] = points[j]

            if (x1 !== x2 && y1 !== y2) {
                if (s.has(`${x1},${y2}`) && s.has(`${x2},${y1}`)) {
                    const area = Math.abs(x2-x1) * Math.abs(y2 - y1)
                    min = Math.min(min, area)
                }
            }
        }
    }

    return min === Infinity ? 0 : min
};