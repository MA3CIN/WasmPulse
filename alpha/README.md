# Running the alpha version
To run the alpha version, first find the PID's of the wasm service that interests you:
```bash
./find_pid.sh
```
Making the script executable will likely be needed. Then, with the chosen PID (1 at a time), run the WasmPulse Resource Measurement Module either via a script:
```bash
./measure.sh <PID>
```
or via a container
```bash
docker run marcinziolkowski:wasmpulse-alpha:latest <PID>
```