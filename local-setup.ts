import {
  packages,
  forPlatform,
  getHostPlatform,
  installTo,
} from "https://deno.land/x/adllang_localsetup@v0.11/mod.ts";

const DENO = packages.deno("1.34.1");
const ADL = packages.adl("1.2");

export async function main() {
  if (Deno.args.length != 1) {
    console.error("Usage: local-setup LOCALDIR");
    Deno.exit(1);
  }
  const localdir = Deno.args[0];

  const platform = getHostPlatform();

  const installs = [
    forPlatform(DENO, platform),
    forPlatform(ADL, platform),
  ];

  await installTo(installs, localdir);
}

main()
  .catch((err) => {
    console.error("error in main", err);
  });
