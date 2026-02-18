class Exo < Formula
  desc "Cloud-Native Bootstrap CLI — scaffold DevOps assets in seconds"
  homepage "https://github.com/Harsh-BH/Exo"
  version "0.1.0"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/Harsh-BH/Exo/releases/download/v#{version}/exo-darwin-arm64"
      sha256 "PLACEHOLDER_DARWIN_ARM64_SHA256"
    else
      url "https://github.com/Harsh-BH/Exo/releases/download/v#{version}/exo-darwin-amd64"
      sha256 "PLACEHOLDER_DARWIN_AMD64_SHA256"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/Harsh-BH/Exo/releases/download/v#{version}/exo-linux-arm64"
      sha256 "PLACEHOLDER_LINUX_ARM64_SHA256"
    else
      url "https://github.com/Harsh-BH/Exo/releases/download/v#{version}/exo-linux-amd64"
      sha256 "PLACEHOLDER_LINUX_AMD64_SHA256"
    end
  end

  def install
    bin.install Dir["exo-*"].first => "exo"
  end

  test do
    assert_match "v#{version}", shell_output("#{bin}/exo version")
  end
end
