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
 * Peer params that might be used due to creation or updation process
 * @export
 * @interface PeerCreateOrUpdateRequest
 */
export interface PeerCreateOrUpdateRequest {
    /**
     * Base64 encoded public key
     * @type {string}
     * @memberof PeerCreateOrUpdateRequest
     */
    publicKey?: any | null;
    /**
     * Base64 encoded preshared key
     * @type {string}
     * @memberof PeerCreateOrUpdateRequest
     */
    presharedKey?: any | null;
    /**
     * Peer's allowed ips, it might be any of IPv4 or IPv6 addresses in CIDR notation
     * @type {Array&lt;string&gt;}
     * @memberof PeerCreateOrUpdateRequest
     */
    allowedIps?: any | null;
    /**
     * Peer's persistend keepalive interval. Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\".
     * @type {string}
     * @memberof PeerCreateOrUpdateRequest
     */
    persistentKeepaliveInterval?: any;
    /**
     * Peer's endpoint in host:port format
     * @type {string}
     * @memberof PeerCreateOrUpdateRequest
     */
    endpoint?: any;
}
