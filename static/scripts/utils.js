function getAverageRGB(imgEl) {
    let blockSize = 5, // only visit every 5 pixels
        defaultRGB = {r: 0, g: 0, b: 0}, // for non-supporting envs
        canvas = document.createElement("canvas"),
        context = canvas.getContext && canvas.getContext("2d"),
        data,
        width,
        height,
        i = -4,
        length,
        rgb = {r: 0, g: 0, b: 0},
        count = 0;

    if (!context) {
        return defaultRGB;
    }

    height = canvas.height =
        imgEl.naturalHeight || imgEl.offsetHeight || imgEl.height;
    width = canvas.width = imgEl.naturalWidth || imgEl.offsetWidth || imgEl.width;

    context.drawImage(imgEl, 0, 0);

    try {
        data = context.getImageData(0, 0, width, height);
    } catch (e) {
        console.log(e);
        /* security error, img on diff domain */ //alert("x");
        return defaultRGB;
    }

    length = data.data.length;

    while ((i += blockSize * 4) < length) {
        ++count;
        rgb.r += data.data[i];
        rgb.g += data.data[i + 1];
        rgb.b += data.data[i + 2];
    }

    // ~~ used to floor values
    rgb.r = ~~(rgb.r / count);
    rgb.g = ~~(rgb.g / count);
    rgb.b = ~~(rgb.b / count);

    return rgb;
}
const LightenColor = function(R, G, B, percent) {
    let amt = Math.round(2.55 * percent);
    R = R + amt;
    B = B + amt;
    G = G + amt;
    return {r: R, g: G, b: B};
};

function openFullscreen() {
    if (elem.requestFullscreen) {
        elem.requestFullscreen();
    } else if (elem.webkitRequestFullscreen) { /* Safari */
        elem.webkitRequestFullscreen();
    } else if (elem.msRequestFullscreen) { /* IE11 */
        elem.msRequestFullscreen();
    }
    document.getElementById("fs-icon").innerHTML = '<i class="fas fa-compress-arrows-alt"></i>';
}

function closeFullscreen() {
    if (document.exitFullscreen) {
        document.exitFullscreen();
    } else if (document.webkitExitFullscreen) { /* Safari */
        document.webkitExitFullscreen();
    } else if (document.msExitFullscreen) { /* IE11 */
        document.msExitFullscreen();
    }
    document.getElementById("fs-icon").innerHTML = '<i class="fas fa-expand-arrows-alt"></i>';
}

function toggleFS() {
    if ((window.fullScreen) ||(window.innerWidth == screen.width && window.innerHeight == screen.height)) {
        closeFullscreen();
    } else {
        openFullscreen();
    }
}

// prevent pinch-to-zoom https://stackoverflow.com/questions/11689353/disable-pinch-zoom-on-mobile-web
document.addEventListener('gesturestart', function (e) {
    e.preventDefault();
    document.body.style.zoom = 1;
});
document.addEventListener('gesturechange', function (e) {
    e.preventDefault();
    document.body.style.zoom = 1;
});
document.addEventListener('gestureend', function (e) {
    e.preventDefault();
    document.body.style.zoom = 1;
});

function isValidHttpUrl(string) {
    let url;

    try {
        url = new URL(string);
    } catch (_) {
        return false;
    }

    return url.protocol === "http:" || url.protocol === "https:";
}