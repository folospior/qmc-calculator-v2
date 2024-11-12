{

    description = "Nix environment for Go";

    inputs.nixpkgs.url = "nixpkgs";

    outputs = { self, nixpkgs, ... }:
        let
            pkgs = nixpkgs.legacyPackages."x86_64-linux";
        in
        {
            devShells."x86_64-linux".default = pkgs.mkShell {
                buildInputs = with pkgs; [ go gopls git direnv gofumpt ];

		        GO111MODULE = "on";
		        GOROOT = pkgs.go;
		        GOPATH = "${pkgs.go}/go";
		        #PATH = "${pkgs.go}/bin:${pkgs.gopls}/bin:${pkgs.gofumpt}/bin";
	        };
        };

	
}
