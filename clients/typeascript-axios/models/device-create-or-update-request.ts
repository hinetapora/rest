/* tslint:disable */
/* eslint-disable */
/**
 * WireGuard RESTful API
 * Manage WireGuard VPN tunnels by RESTful manner.  Supported features:    * Manage device: create, update, and delete wireguard interface   * Manage device's ip addresses: attache or detach ip addresses to the netowrk interface   * Manage device's peers: create, update, and delete peers   * Peer's QR code, for use in WireGuard & ForestVPN client  ForestVPN client may be used as alternative client with enabled P2P technology over WireGuard tunnelling. Read more on https://forestvpn.com/ 
 *
 * OpenAPI spec version: 1.0
 * Contact: support@forestvpn.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */
/**
 * Device params that might be used due to creation or updation process
 * @export
 * @interface DeviceCreateOrUpdateRequest
 */
export interface DeviceCreateOrUpdateRequest {
    /**
     * WireGuard device name. Usually it is network interface name
     * @type {string}
     * @memberof DeviceCreateOrUpdateRequest
     */
    name?: any | null;
    /**
     * WireGuard device listen port.
     * @type {number}
     * @memberof DeviceCreateOrUpdateRequest
     */
    listenPort?: any | null;
    /**
     * WireGuard device private key encoded by base64.
     * @type {string}
     * @memberof DeviceCreateOrUpdateRequest
     */
    privateKey?: any | null;
    /**
     * WireGuard device firewall mark.
     * @type {number}
     * @memberof DeviceCreateOrUpdateRequest
     */
    firewallMark?: any | null;
    /**
     * IPv4 or IPv6 addresses in CIDR notation
     * @type {Array&lt;string&gt;}
     * @memberof DeviceCreateOrUpdateRequest
     */
    networks?: any | null;
}
