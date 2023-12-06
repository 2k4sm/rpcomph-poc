import bluetooth
import time

def discover_devices(timeout=10):
    print("Scanning for Bluetooth devices (Timeout: {} seconds)...".format(timeout))
    
    start_time = time.time()
    unique_devices = set()

    while time.time() - start_time < timeout:
        nearby_devices = bluetooth.discover_devices(duration=8, lookup_names=True, lookup_class=True, device_id=-1,
                                                    device_name="", device_class=0, service_classes=None,
                                                    device_oui=None, lookup_oui=True, lookup_oui_in_manufacturer=True)

        for addr, name, _ in nearby_devices:
            if (name, addr) not in unique_devices:
                unique_devices.add((name, addr))

    print("Scanning complete. Found {} unique devices.".format(len(unique_devices)))

    for name, addr in unique_devices:
        print("Device Name: {}, Address: {}".format(name, addr))

if __name__ == "__main__":
    discover_devices()
