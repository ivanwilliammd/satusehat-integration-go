# Changelog

## Unreleased

### Added

- **HL7 v2 + MLLP** (`src/hl7v2`): `Parse`/`Message`/`Segment` (ER7, path access, escape decode and
  encode, byte-for-byte round trip), MLLP framing with a stateful `Reader`, `BuildAck`/`BuildNak`
  (AA/AE/AR with MSA-2 echo), `Client` with `Send` and `Session` (one connection, one result each),
  and declarative `Profile`s (`AdtA01`, `AdtA08`, `OruR01`).
- **Shared corpus**: the 20 fixtures in `src/hl7v2/testdata` are copied byte-for-byte from the PHP
  SDK, so parity is tested with the same messages instead of being re-described.

### Fixed

- `Shift` panicked when a frame's trailing CR had not arrived yet (`slice bounds out of range`); the
  remainder is now empty instead, so the reader simply waits for the rest of the frame.
