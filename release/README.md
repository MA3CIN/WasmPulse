# Running the release version
To run the alpha version, first find the PID's of the wasm service that interests you:
```bash
./find_pid.sh
```
Making the script executable will likely be needed. Then, with the chosen PID (1 at a time), run the WasmPulse platform with Docker (or container runtime of your choice):
```bash
docker run marcinziolkowski:wasmpulse:latest
```

The WasmPulse platform will automatically find all PID's of your Wasm services, and allow you to monitor their resource usage in real time.

Visit the API on https://localhost:8080 and acess the relevant data for your service of choice.