const BACK_IMAGE = "/img/back.png";
const FOLDER_IMAGE = "/img/folder.png";

function render(data) {
    if (data != undefined && data.success) {
        if (data.success) {
            console.log(data)
            data.message.forEach(element => {
                let logo = "/img/folder.png";
                if (element.logo != "") {
                    if (isValidHttpUrl(element.logo)) {
                        logo = element.logo;
                    } else {
                        logo = `/icons/${element.logo}`;
                    }
                }
                let elementId = `${element.name}`;
                if (element.is_folder === true) {
                    console.log("folder")
                    logo = FOLDER_IMAGE;
                    elementId = "folder-" + elementId;
                }
                
                document.getElementById("deck").innerHTML += `
                                    <div class="element" id="${elementId}" onclick="executeCommand(this.id)">
                                        <div style="text-align: center;">
                                            <img src="${logo}" id="img-${element.name}" crossorigin="*" draggable="false" />
                                        </div>
                                    </div>
                                `;
                let aa = document.getElementById("img-" + element.name);
                aa.addEventListener("load", () => {
                    let c = getAverageRGB(aa);
                    let luma = 0.2126 * c.r + 0.7152 * c.g + 0.0722 * c.b; // per ITU-R BT.709
                    console.log("LUMA " + luma + "  " + element.name);
                    if (luma < 40) {
                        c = LightenColor(c.r, c.g, c.b, +10);
                        console.log("!! COLOR TOO DARK !! " + element.name)
                    } else {
                        c = LightenColor(c.r, c.g, c.b, -10);
                    }
                    document.getElementById(elementId).style.backgroundColor = `rgb(${Math.abs(c.r)}, ${Math.abs(c.g)}, ${Math.abs(c.b)})`;
                });
            })
        } else {
            console.log(`Something went wrong`);
        }
    }
}