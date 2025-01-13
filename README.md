# enigma

example of .env file


ENIGMA_REFLECTOR=1
ENIGMA_KEY=1-0 2-10 3-0
ENIGMA_MESSAGE=HELLO

ENIGMA_REFLECTOR

Description: Specifies the reflector to be used in the Enigma machine. The reflector determines how signals are bounced back through the rotors.

ENIGMA_KEY

Description: Defines the rotor configuration and their initial positions for the Enigma machine.
Format: Space-separated rotor configurations, where each rotor configuration is in the format<rotor_number>-<initial_offset>

Documentation for Enigma Configuration Environment File

The environment file configures the settings for an Enigma machine simulation. Below is a detailed explanation of each parameter:
ENIGMA_REFLECTOR

    Description: Specifies the reflector to be used in the Enigma machine. The reflector determines how signals are bounced back through the rotors.
    Value Type: Integer
    Example Values:
        1: Indicates the use of the first reflector configuration.
        2: Indicates an alternative reflector (if implemented).
    Default Behavior: If not set, the program may default to a predefined reflector.

ENIGMA_KEY

    Description: Defines the rotor configuration and their initial positions for the Enigma machine.
    Format: Space-separated rotor configurations, where each rotor configuration is in the format <rotor_number>-<initial_offset>.
        <rotor_number>: The rotor's identifier (e.g., 1, 2, 3).
        <initial_offset>: The starting position of the rotor, specified as an integer (e.g., 0 for no offset, 10 for a shift of 10).
    Value Type: String
    Example Value:
        1-0 2-10 3-0:
            Rotor 1 starts at position 0.
            Rotor 2 starts at position 10.
            Rotor 3 starts at position 0.
    Notes: Ensure that the rotor numbers and offsets match the rotors configured in the program.

ENIGMA_MESSAGE

Description: The plaintext message to be encrypted using the Enigma machine.

To run code:
```bash
go run main.go
```