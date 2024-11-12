{

    description = "Nix environment for Go";

    inputs.nixpkgs.url = "nixpkgs";

    outputs = { self, nixpkgs, ... }:
        let
            pkgs = nixpkgs.legacyPackages.${builtins.currentSystem};
        in
        {
            devShell = pkgs.mkShell {
                buildInputs = with pkgs; [ go gopls git direnv gofumpt ];

		GO111MODULE = "on";
		GOROOT = pkgs.go;
		GOPATH = "${pkgs.go}/go";
		PATH = "${pkgs.go}/bin:${pkgs.gopls}/bin:${pkgs.gofumpt}/bin";
	    };
        };
	
}
