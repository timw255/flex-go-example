const gulp = require('gulp');
const del = require('del');
const fs = require('fs');
const zlib = require('zlib');
const path = require('path');

gulp.task('clean', function () {
    return del('dist');
});

gulp.task('dist', (done) => {
    gulp.src(['index.js', 'package.json'])
        .pipe(gulp.dest('dist'))
        .on('finish', (err) => {
            if (err) { throw error; }
            const zip = zlib.createGzip();
            const fileContents = fs.createReadStream(path.resolve(__dirname, 'main.wasm'));
            const writeStream = fs.createWriteStream(path.resolve(__dirname, 'dist', 'main.wasm.gz'));

            fileContents
                .pipe(zip)
                .pipe(writeStream)
                .on('finish', (err) => {
                    if (err) { throw error; }
                    done();
                });
        });
});

gulp.task('build-dist', gulp.series('clean', 'dist'));

gulp.task('default', gulp.series('build-dist'));