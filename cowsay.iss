[Setup]
AppName=cowsay
AppVersion=1.0
DefaultDirName={pf64}\Cowsay

[Files]
Source: "cowsay.exe"; DestDir: "{app}"

[Registry]
Root: HKLM; Subkey: "SYSTEM\CurrentControlSet\Control\Session Manager\Environment";ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{pf64}\Cowsay";Check: NeedsAddPath('{pf64}\Cowsay')

[Code]

function NeedsAddPath(Param: string): boolean;
var
  OrigPath: string;
begin
  if not RegQueryStringValue(HKEY_LOCAL_MACHINE,
    'SYSTEM\CurrentControlSet\Control\Session Manager\Environment',
    'Path', OrigPath) 
  then begin
    Result := True;
    exit;
  end;
  // look for the path with leading and trailing semicolon
  // Pos() returns 0 if not found
  Result := Pos(';' + Param + ';', ';' + OrigPath + ';') = 0;
end;