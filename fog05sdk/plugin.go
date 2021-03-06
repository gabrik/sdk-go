package fog05sdk

import (
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/google/uuid"
)

// OS is the object to interact with OS Plugin
type OS struct {
	uuid      string
	connector *YaksConnector
	node      string
}

// CallOSPluginFunction calls an Eval registered within the OS Plugin, returns a pointer to a genering interface{}
func (os *OS) CallOSPluginFunction(fname string, fparameters map[string]interface{}) (*string, error) {
	res, err := os.connector.Local.Actual.ExecOSEval(os.node, fname, fparameters)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		er := FError{*res.ErrorMessage + " ErrNo: " + string(*res.Error), nil}
		return nil, &er
	}
	return res.Result, nil
}

// DirExists check if the given directory exists
func (os *OS) DirExists(dirpath string) (bool, error) {
	r, err := os.CallOSPluginFunction("dir_exists", map[string]interface{}{"dir_path": dirpath})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// CreateDir creates the given directory
func (os *OS) CreateDir(dirpath string) (bool, error) {
	r, err := os.CallOSPluginFunction("create_dir", map[string]interface{}{"dir_path": dirpath})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// RemoveDir removes the given directory
func (os *OS) RemoveDir(dirpath string) (bool, error) {
	r, err := os.CallOSPluginFunction("remove_dir", map[string]interface{}{"dir_path": dirpath})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// DownloadFile downloads the given file into the given path
func (os *OS) DownloadFile(url string, filepath string) (bool, error) {
	r, err := os.CallOSPluginFunction("download_file", map[string]interface{}{"url": url, "file_path": filepath})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// ExecuteCommand executes the given command, with given flags
func (os *OS) ExecuteCommand(command string, blocking bool, external bool) (string, error) {
	r, err := os.CallOSPluginFunction("execute_command", map[string]interface{}{"command": command, "blocking": blocking, "external": external})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// CreateFile creates the empty given file
func (os *OS) CreateFile(filepath string) (bool, error) {
	r, err := os.CallOSPluginFunction("create_file", map[string]interface{}{"file_path": filepath})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// RemoveFile removes the given file
func (os *OS) RemoveFile(filepath string) (bool, error) {
	r, err := os.CallOSPluginFunction("remove_file", map[string]interface{}{"file_path": filepath})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// StoreFile creates and stores the given content into the given file
func (os *OS) StoreFile(content string, filepath string, filename string) (bool, error) {

	c := hex.EncodeToString([]byte(b64.StdEncoding.EncodeToString([]byte(content))))
	r, err := os.CallOSPluginFunction("store_file", map[string]interface{}{"file_path": filepath, "filename": filename, "content": c})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// ReadFile reads the given file
func (os *OS) ReadFile(filepath string, root bool) (string, error) {
	r, err := os.CallOSPluginFunction("read_file", map[string]interface{}{"file_path": filepath, "root": root})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// FileExists checks if the given file exists
func (os *OS) FileExists(filepath string) (bool, error) {
	r, err := os.CallOSPluginFunction("file_exists", map[string]interface{}{"file_path": filepath})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// SendSigInt sends INT signal to the given PID
func (os *OS) SendSigInt(pid int) (bool, error) {
	r, err := os.CallOSPluginFunction("send_sig_int", map[string]interface{}{"pid": pid})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// SendSigKill sends the KILL signal to the given PID
func (os *OS) SendSigKill(pid int) (bool, error) {
	r, err := os.CallOSPluginFunction("send_sig_kill", map[string]interface{}{"pid": pid})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// CheckIfPIDExists check if the PID is still running
func (os *OS) CheckIfPIDExists(pid int) (bool, error) {
	r, err := os.CallOSPluginFunction("check_if_pid_exists", map[string]interface{}{"pid": pid})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// GetInterfaceType get the interface type for the given network interface
func (os *OS) GetInterfaceType(facename string) (string, error) {
	r, err := os.CallOSPluginFunction("get_intf_type", map[string]interface{}{"name": facename})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// SetInterfaceUnaviable sets the given network interface as unaviable
func (os *OS) SetInterfaceUnaviable(facename string) (bool, error) {
	r, err := os.CallOSPluginFunction("set_interface_unaviable", map[string]interface{}{"intf_name": facename})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// SetInterfaceAvailable sets the given network interface as available
func (os *OS) SetInterfaceAvailable(facename string) (bool, error) {
	r, err := os.CallOSPluginFunction("set_interface_available", map[string]interface{}{"intf_name": facename})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// Checksum computes the checksum (SHA256) for the given file
func (os *OS) Checksum(filepath string) (string, error) {
	r, err := os.CallOSPluginFunction("checksum", map[string]interface{}{"file_path": filepath})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// LocalMgmtAddress gets the local management ip address
func (os *OS) LocalMgmtAddress() (string, error) {
	r, err := os.CallOSPluginFunction("local_mgmt_address", map[string]interface{}{})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// NM is the object to interact with the network manager plugin
type NM struct {
	uuid      string
	connector *YaksConnector
	node      string
}

// CallNMPluginFunction calls an Eval register within the network manager, returns a genering pointer to interface{}
func (nm *NM) CallNMPluginFunction(fname string, fparameters map[string]interface{}) (*string, error) {
	res, err := nm.connector.Local.Actual.ExecNMEval(nm.node, nm.uuid, fname, fparameters)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		er := FError{*res.ErrorMessage + " ErrNo: " + string(*res.Error), nil}
		return nil, &er
	}
	return res.Result, nil
}

// CreateVirtualInterface creates the given virtual interface and returns its information
func (nm *NM) CreateVirtualInterface(intfid string, descriptor FDUInterfaceRecord) (*map[string]interface{}, error) {

	jd, err := json.Marshal(descriptor)
	var md map[string]interface{}

	json.Unmarshal(jd, &md)

	r, err := nm.CallNMPluginFunction("create_virtual_interface", map[string]interface{}{"intf_id": intfid, "descriptor": md})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// DeleteVirtualInterface deletes the given network interface and returns its information
func (nm *NM) DeleteVirtualInterface(intfid string) (*string, error) {
	r, err := nm.CallNMPluginFunction("delete_virtual_interface", map[string]interface{}{"intf_id": intfid})
	if err != nil {
		return nil, err
	}

	return r, nil
}

// CreateVirtualBridge creates the given virtual bridge and returns its information
func (nm *NM) CreateVirtualBridge(name string, uuid string) (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("create_virtual_bridge", map[string]interface{}{"name": name, "uuid": uuid})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// DeleteVirtualBridge removes the given virtual bridge and returns its information
func (nm *NM) DeleteVirtualBridge(uuid string) (string, error) {
	r, err := nm.CallNMPluginFunction("delete_virtual_bridge", map[string]interface{}{"br_uuid": uuid})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// CreateBridgesIfNotExists create the given bridges if they are not existing and returns a slice with the bridges informations
func (nm *NM) CreateBridgesIfNotExists(expected []string) (*[]map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("create_bridges_if_not_exist", map[string]interface{}{"expected_bridges": expected})
	if err != nil {
		return nil, err
	}

	myVar := [](map[string]interface{}){}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// ConnectInterfaceToConnectionPoint connects the given interface to the given connection point and returns interface information
func (nm *NM) ConnectInterfaceToConnectionPoint(intfid string, cpid string) (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("connect_interface_to_connection_point", map[string]interface{}{"intf_id": intfid, "cp_id": cpid})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// DisconnectInterface disconnects the given interface and returns its information
func (nm *NM) DisconnectInterface(intfid string) (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("disconnect_interface", map[string]interface{}{"intf_id": intfid})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// ConnectCPToVNetwork connect the given connection point to the given network and returns connection point information
func (nm *NM) ConnectCPToVNetwork(cpid string, vnetid string) (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("connect_cp_to_vnetwork", map[string]interface{}{"cp_id": cpid, "vnet_id": vnetid})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// DisconnectCP disconnect the given connection point and returns its information
func (nm *NM) DisconnectCP(cpid string) (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("disconnect_cp", map[string]interface{}{"cp_id": cpid})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// DeletePort deletes the given connection point
func (nm *NM) DeletePort(cpid string) (bool, error) {
	r, err := nm.CallNMPluginFunction("delete_port", map[string]interface{}{"cp_id": cpid})
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(*r)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return false, &er
	}
	return b, nil
}

// GetAddress gets the IP address of the specified connection point
func (nm *NM) GetAddress(cpid string) (string, error) {
	r, err := nm.CallNMPluginFunction("delete_port", map[string]interface{}{"cp_id": cpid})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// AddPortToRouter adds the given port to the given router and returns router information
func (nm *NM) AddPortToRouter(routerid string, porttype string, vnetid string, ipaddress string) (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("add_router_port", map[string]interface{}{"router_id": routerid, "port_type": porttype, "vnet_id": vnetid, "ip_address": ipaddress})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// RemovePortFromRouter remove the given port from the given router and returns router information
func (nm *NM) RemovePortFromRouter(routerid string, vnetid string) (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("remove_port_from_router", map[string]interface{}{"router_id": routerid, "vnet_id": vnetid})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// CreateFloatingIP creates a floating IP and returns its information
func (nm *NM) CreateFloatingIP() (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("create_floating_ip", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// DeleteFloatingIP deletes the given floaing IP and returns its information
func (nm *NM) DeleteFloatingIP(ipid string) (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("delete_floating_ip", map[string]interface{}{"ip_id": ipid})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// AssignFloatingIP assign the given floating IP to the given connection point and returns floating IP information
func (nm *NM) AssignFloatingIP(ipid string, cpid string) (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("assign_floating_ip", map[string]interface{}{"ip_id": ipid, "cp_id": cpid})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// RemoveFloatingIP retain the given floating ip from the given connection point and returns floating IP information
func (nm *NM) RemoveFloatingIP(ipid string, cpid string) (*map[string]interface{}, error) {
	r, err := nm.CallNMPluginFunction("remove_floating_ip", map[string]interface{}{"ip_id": ipid, "cp_id": cpid})
	if err != nil {
		return nil, err
	}

	myVar := make(map[string]interface{})
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// GetOverlayFace gets the configured network interface for overlay networks
func (nm *NM) GetOverlayFace() (string, error) {
	r, err := nm.CallNMPluginFunction("get_overlay_face", map[string]interface{}{})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// GetVLANFace gets the configured network interfaces for VLAN networks
func (nm *NM) GetVLANFace() (string, error) {
	r, err := nm.CallNMPluginFunction("get_vlan_face", map[string]interface{}{})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// AddNodePort creates a new network port in the node
func (nm *NM) AddNodePort(cp ConnectionPointRecord) error {
	return nm.connector.Local.Desired.AddNodePort(nm.node, nm.uuid, cp.UUID, cp)
}

// GetNodePort gets the given port information
func (nm *NM) GetNodePort(cpid string) (*ConnectionPointRecord, error) {
	return nm.connector.Local.Desired.GetNodePort(nm.node, nm.uuid, cpid)
}

// GetAllNodePorts gets information about all the port in the node
func (nm *NM) GetAllNodePorts() ([]ConnectionPointRecord, error) {
	return nm.connector.Local.Desired.GetAllNodePorts(nm.node, nm.uuid)
}

// RemoveNodePort removes the given port
func (nm *NM) RemoveNodePort(cpid string) error {

	cpd, _ := nm.GetNodePort(cpid)
	(*cpd).Status = DESTROY

	return nm.connector.Local.Desired.AddNodePort(nm.node, nm.uuid, cpid, *cpd)
}

// CreateConnectionPoint creates the given connection point
func (nm *NM) CreateConnectionPoint(descriptor ConnectionPointDescriptor) (*ConnectionPointRecord, error) {
	v, err := json.Marshal(descriptor)
	if err != nil {
		return nil, err
	}
	var md map[string]interface{}

	err = json.Unmarshal(v, &md)
	if err != nil {
		return nil, err
	}

	r, err := nm.CallNMPluginFunction("create_port_agent", map[string]interface{}{"descriptor": md})
	if err != nil {
		return nil, err
	}

	myVar := ConnectionPointRecord{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil

}

// RemoveConnectionPoint removes the given connection point
func (nm *NM) RemoveConnectionPoint(cpid string) (*ConnectionPointRecord, error) {
	r, err := nm.CallNMPluginFunction("destroy_port_agent", map[string]interface{}{"cp_id": cpid})
	if err != nil {
		return nil, err
	}

	myVar := ConnectionPointRecord{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// CreateMACVLANInterface creates a MACVLAN interface over the given interface
func (nm *NM) CreateMACVLANInterface(masterIntf string) (string, error) {
	r, err := nm.CallNMPluginFunction("create_macvlan_interface", map[string]interface{}{"master_intf": masterIntf})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// DeleteMACVLANInterface deletes the given MACVLAN interface
func (nm *NM) DeleteMACVLANInterface(intfName string, netns string) (string, error) {
	if netns == "" {
		netns = "1"
	}
	r, err := nm.CallNMPluginFunction("delete_macvlan_interface", map[string]interface{}{"intfName": intfName, "netns": netns})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// CreateNetworkNamespace creates a new network namespace, and returns its name
func (nm *NM) CreateNetworkNamespace() (string, error) {
	r, err := nm.CallNMPluginFunction("create_network_namespace", map[string]interface{}{})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// DeleteNetworkNamespace deletes the given network namespace, and returns its name
func (nm *NM) DeleteNetworkNamespace(netns string) (string, error) {
	r, err := nm.CallNMPluginFunction("delete_network_namespace", map[string]interface{}{"nsname": netns})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// MoveInterfaceInNamespace moves the given interface to the given namespace, is netns is empty will move to the default namespace
func (nm *NM) MoveInterfaceInNamespace(intfName string, netns string) (*InterfaceInfo, error) {
	if netns == "" {
		netns = "1"
	}
	r, err := nm.CallNMPluginFunction("move_interface_in_namespace", map[string]interface{}{"intf_name": intfName, "nsname": netns})
	if err != nil {
		return nil, err
	}

	myVar := InterfaceInfo{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// RenameVirtualInterfaceInNamespace renames the given interface
func (nm *NM) RenameVirtualInterfaceInNamespace(name string, newname string, nsname string) (string, error) {
	var r *string
	var err error
	if nsname == "" {
		r, err = nm.CallNMPluginFunction("rename_virtual_interface_in_namespace", map[string]interface{}{"name": name, "newname": newname})
	} else {
		r, err = nm.CallNMPluginFunction("rename_virtual_interface_in_namespace", map[string]interface{}{"name": name, "newname": newname, "nsname": nsname})
	}
	if err != nil {
		return "", err
	}

	return *r, nil
}

// AttachInterfaceToBridge attaches the given interface to the given bridge
func (nm *NM) AttachInterfaceToBridge(intfName string, brName string) (*InterfaceInfo, error) {
	r, err := nm.CallNMPluginFunction("attach_interface_to_bridge", map[string]interface{}{"intf_name": intfName, "br_name": brName})
	if err != nil {
		return nil, err
	}

	myVar := InterfaceInfo{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// DetachInterfaceFromBridge detaches the interface from the current connected bridge
func (nm *NM) DetachInterfaceFromBridge(intfName string) (*InterfaceInfo, error) {
	r, err := nm.CallNMPluginFunction("detach_interface_from_bridge", map[string]interface{}{"intf_name": intfName})
	if err != nil {
		return nil, err
	}

	myVar := InterfaceInfo{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// CreateVirtualInterfaceInNamespace creates a veth pair in the given network namespace, with the given name for the internal interface
func (nm *NM) CreateVirtualInterfaceInNamespace(intfName string, netns string) (*NamespaceInfo, error) {
	r, err := nm.CallNMPluginFunction("create_virtual_interface_in_namespace", map[string]interface{}{"internal_name": intfName, "nsname": netns})
	if err != nil {
		return nil, err
	}

	myVar := NamespaceInfo{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// DeleteVirtualInterfaceFromNamespace deletes the given interface from the the given network namespace
func (nm *NM) DeleteVirtualInterfaceFromNamespace(intfName string, netns string) (*NamespaceInfo, error) {
	r, err := nm.CallNMPluginFunction("delete_virtual_interface_from_namespace", map[string]interface{}{"internal_name": intfName, "nsname": netns})
	if err != nil {
		return nil, err
	}

	myVar := NamespaceInfo{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// AssignAddressToInterfaceInNamespace assigns the given address to the given interface in the the given network namespace, address are in the form AAA.AAA.AAA.AAA/NM
func (nm *NM) AssignAddressToInterfaceInNamespace(intfName string, netns string, address string) (*NamespaceInfo, error) {
	var r *string
	var err error
	if address == "" {
		r, err = nm.CallNMPluginFunction("assign_address_to_interface_in_namespace", map[string]interface{}{"intf_name": intfName, "nsname": netns})
	} else {
		r, err = nm.CallNMPluginFunction("assign_address_to_interface_in_namespace", map[string]interface{}{"intf_name": intfName, "nsname": netns, "address": address})
	}
	if err != nil {
		return nil, err
	}

	myVar := NamespaceInfo{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// AssignMACAddressToInterfaceInNamespace assigns the given address to the given interface in the the given network namespace, address are in the form AA:BB:CC:DD:EE:FF
func (nm *NM) AssignMACAddressToInterfaceInNamespace(intfName string, netns string, address string) (*NamespaceInfo, error) {
	r, err := nm.CallNMPluginFunction("assign_mac_address_to_interface_in_namespace", map[string]interface{}{"intf_name": intfName, "nsname": netns, "address": address})
	if err != nil {
		return nil, err
	}

	myVar := NamespaceInfo{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// GetAddressOfInterfaceInNamespace retrieves the address to the given interface in the the given network namespace
func (nm *NM) GetAddressOfInterfaceInNamespace(intfName string, netns string) (*InterfaceInfo, error) {
	r, err := nm.CallNMPluginFunction("get_address_of_interface_in_namespace", map[string]interface{}{"intf_name": intfName, "nsname": netns})
	if err != nil {
		return nil, err
	}

	myVar := InterfaceInfo{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// RemoveAddressFromInterfaceInNamespace removes the address from the given interface in the the given network namespace
func (nm *NM) RemoveAddressFromInterfaceInNamespace(intfName string, netns string) (*NamespaceInfo, error) {
	r, err := nm.CallNMPluginFunction("remove_address_from_interface_in_namespace", map[string]interface{}{"intf_name": intfName, "nsname": netns})
	if err != nil {
		return nil, err
	}

	myVar := NamespaceInfo{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// Agent is the object to interect with the Agent
type Agent struct {
	connector *YaksConnector
	node      string
}

// CallAgentFunction calls an Eval registered within the Agent and returns a generic pointer to interface
func (ag *Agent) CallAgentFunction(fname string, fparameters map[string]interface{}) (*string, error) {
	res, err := ag.connector.Local.Actual.ExecAgentEval(ag.node, fname, fparameters)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		er := FError{*res.ErrorMessage + " ErrNo: " + string(*res.Error), nil}
		return nil, &er
	}
	return res.Result, nil
}

// GetImageInfo given an image UUID retruns the FDUImage object associated
func (ag *Agent) GetImageInfo(imgid string) (*FDUImage, error) {
	r, err := ag.CallAgentFunction("get_image_info", map[string]interface{}{"image_uuid": imgid})
	if err != nil {
		return nil, err
	}

	myVar := FDUImage{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// GetFDUInfo given a node id, fdu id and instance id returns the FDU object associated
func (ag *Agent) GetFDUInfo(nodeid string, fduid string, instanceid string) (*FDU, error) {
	r, err := ag.CallAgentFunction("get_node_fdu_info", map[string]interface{}{"fdu_uuid": fduid, "instance_uuid": instanceid, "node_uuid": nodeid})
	if err != nil {
		return nil, err
	}

	myVar := FDU{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// GetFDUDescriptor returns the descriptor for the given FDU ID
func (ag *Agent) GetFDUDescriptor(fduid string) (*FDU, error) {
	r, err := ag.CallAgentFunction("get_fdu_info", map[string]interface{}{"fdu_uuid": fduid})
	if err != nil {
		return nil, err
	}

	myVar := FDU{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// GetNetworkInfo given a network id returns the VirtualNetwork object associated
func (ag *Agent) GetNetworkInfo(netid string) (*VirtualNetwork, error) {
	r, err := ag.CallAgentFunction("get_network_info", map[string]interface{}{"uuid": netid})
	if err != nil {
		return nil, err
	}

	myVar := VirtualNetwork{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// GetPortInfo given a connection point id returns the ConnectionPointDescriptor associated
func (ag *Agent) GetPortInfo(cpid string) (*ConnectionPointDescriptor, error) {
	r, err := ag.CallAgentFunction("get_port_info", map[string]interface{}{"cp_uuid": cpid})
	if err != nil {
		return nil, err
	}

	myVar := ConnectionPointDescriptor{}
	err = json.Unmarshal([]byte(*r), &myVar)
	if err != nil {
		er := FError{"Error on conversion: " + err.Error(), nil}
		return nil, &er
	}

	return &myVar, nil
}

// GetNodeMGMTAddress given a node id return the node management IP address
func (ag *Agent) GetNodeMGMTAddress(nodeid string) (string, error) {
	r, err := ag.CallAgentFunction("get_node_mgmt_address", map[string]interface{}{"node_uuid": nodeid})
	if err != nil {
		return "", err
	}

	return *r, nil
}

// FOSPlugin rapresents an Eclipse fog05 Plugin
type FOSPlugin struct {
	version   int
	connector *YaksConnector
	node      string
	NM        *NM
	OS        *OS
	Agent     *Agent
	UUID      string
}

// NewPlugin returns a new FOSPlugin object
func NewPlugin(version int, pluginuuid string) *FOSPlugin {
	if pluginuuid == "" {
		pluginuuid = uuid.UUID.String(uuid.New())
	}
	return &FOSPlugin{version: version, UUID: pluginuuid, node: "", NM: nil, OS: nil, connector: nil, Agent: nil}
}

// GetOSPlugin loads the OS plugin discovering it from YAKS
func (pl *FOSPlugin) GetOSPlugin() bool {
	pls, err := pl.connector.Local.Actual.GetAllPlugins(pl.node)
	if err != nil {
		panic(err.Error())
	}
	for _, pid := range pls {
		pld, err := pl.connector.Local.Actual.GetNodePlugin(pl.node, pid)
		if err != nil {
			panic(err.Error())
		}
		if pld.Type == "os" {
			pl.OS = &OS{uuid: pld.UUID, connector: pl.connector, node: pl.node}
			return true
		}
	}
	return false
}

// GetNMPlugin loads the Network Manager plugin discovering it from YAKS
func (pl *FOSPlugin) GetNMPlugin() bool {
	pls, err := pl.connector.Local.Actual.GetAllPlugins(pl.node)
	if err != nil {
		panic(err.Error())
	}
	for _, pid := range pls {
		pld, err := pl.connector.Local.Actual.GetNodePlugin(pl.node, pid)
		if err != nil {
			panic(err.Error())
		}
		if pld.Type == "network" {
			pl.NM = &NM{uuid: pld.UUID, connector: pl.connector, node: pl.node}
			return true
		}
	}
	return false
}

// GetAgent loads the Agent discovering it from YAKS
func (pl *FOSPlugin) GetAgent() bool {
	pls, err := pl.connector.Local.Actual.GetAllPlugins(pl.node)
	if err != nil {
		panic(err.Error())
	}
	for _, pid := range pls {
		pld, err := pl.connector.Local.Actual.GetNodePlugin(pl.node, pid)
		if err != nil {
			panic(err.Error())
		}
		if pld.Type == "network" {
			pl.Agent = &Agent{connector: pl.connector, node: pl.node}
			return true
		}
	}
	return false
}

// GetLocalMGMTAddress returns the local management IP address
func (pl *FOSPlugin) GetLocalMGMTAddress() string {
	ip, err := pl.OS.LocalMgmtAddress()
	if err != nil {
		panic(err.Error())
	}
	return ip
}

// GetNodeConfiguration returns the node configuration
func (pl *FOSPlugin) GetNodeConfiguration() (*NodeConfiguration, error) {
	c, err := pl.connector.Local.Actual.GetNodeConfiguration(pl.node)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// GetPluginState returns the plugin state, retrives it from YAKS, as a map[string]interface, each implementation of the plugin can have his own state representation
func (pl *FOSPlugin) GetPluginState() map[string]interface{} {
	s, err := pl.connector.Local.Actual.GetNodePluginState(pl.node, pl.UUID)
	if err != nil {
		panic(err.Error())
	}
	return *s
}

// SavePluginState stores the plugin state into YAKS
func (pl *FOSPlugin) SavePluginState(state map[string]interface{}) error {
	return pl.connector.Local.Actual.AddNodePluginState(pl.node, pl.UUID, state)
}

// RemovePluginState removes the plugin state from YAKS
func (pl *FOSPlugin) RemovePluginState() error {
	return pl.connector.Local.Actual.RemoveNodePluginState(pl.node, pl.UUID)
}
