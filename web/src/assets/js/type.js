export class User {
    constructor(phone, area, password) {
        this.phone = phone;
        this.area = area;
        this.password = password;
    }
}

export class Order {
    constructor(start, dest) {
        this.start = start;
        this.dest = dest;
    }
}

export class Trace {
    constructor(oid, current, next) {
        this.oid = oid;
        this.current = current;
        this.next = next;
    }
}

export const Error = {
    10001: "参数错误",
    10002: "注册失败",
    10003: "登录失败",
    10004: "Cookies不合法",
    10005: "用户未登录",
    10006: "手机号不合法",
    10007: "用户已注册",
    10008: "无法设置Cookies，请检查浏览器设置",
    10009: "地理位置时效，请重新登录",
    10010: "创建订单失败",
    10011: "获取订单失败",
    10012: "获取订单信息失败",
    10013: "创建路径失败",
    10014: "无法获取当前位置",
    10015: "完成订单失败",
    10016: "地址不合法",
    10017: "获取路径失败",
};