function ConfigMgr() {
    this.base_address = "http://127.0.0.1:8080"
}

ConfigMgr.prototype.get_address_config = function (target) {
    let url = ""
    switch (target) {
        case "GET_ALL_PREPARES":
            url = this.base_address + ""
            break;
        default:
            console.log("get_address_config, invalid target: " + target)
            break;
    }
    return url
}

export default ConfigMgr
