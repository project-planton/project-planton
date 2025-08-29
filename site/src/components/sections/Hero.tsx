"use client";
import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Copy, Check, Github, ArrowDown } from "lucide-react";

export default function Hero() {
	const [copied, setCopied] = useState(false);
	const installCommand = "brew install project-planton/tap/project-planton";

	const copyToClipboard = () => {
		navigator.clipboard.writeText(installCommand);
		setCopied(true);
		setTimeout(() => setCopied(false), 2000);
	};

	const scrollToExample = () => {
		const element = document.getElementById("examples");
		if (element) {
			element.scrollIntoView({ behavior: "smooth" });
		}
	};

	return (
		<div className="relative overflow-hidden font-sans">
			{/* Deep purple gradient to match production */}
			<div className="absolute inset-0 bg-gradient-to-b from-[#0b0713] via-[#140a26] to-[#0a0912]" />
			{/* Radial magenta glow */}
			<div className="absolute inset-0 pointer-events-none bg-[radial-gradient(1200px_500px_at_50%_8%,rgba(217,70,239,0.18),rgba(0,0,0,0))]" />

			<div className="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
				<div className="min-h-[calc(100vh-64px)] flex items-center justify-center">
					<div className="text-center max-w-4xl mx-auto">
						{/* Badges */}
						<div className="flex flex-wrap justify-center gap-2 mb-10">
							<Badge variant="outline" className="rounded-full border-white/15 text-white/80 bg-white/10">Apache-2.0</Badge>
							<Badge variant="outline" className="rounded-full border-white/15 text-white/80 bg-white/10">KRM/Protobuf/Buf</Badge>
							<Badge variant="outline" className="rounded-full border-white/15 text-white/80 bg-white/10">Pulumi/OpenTofu</Badge>
							<Badge variant="outline" className="rounded-full border-white/15 text-white/80 bg-white/10">CLI-first</Badge>
							<Badge variant="outline" className="rounded-full border-white/15 text-white/80 bg-white/10">CI/CD-ready</Badge>
						</div>

						{/* Heading */}
						<h1 className="text-4xl sm:text-6xl lg:text-7xl font-extrabold tracking-tight mb-6">
							<span className="text-white">Open‑Source Multi‑Cloud</span>
							<br />
							<span className="bg-gradient-to-r from-[#f0abfc] via-[#f472b6] to-[#d946ef] bg-clip-text text-transparent">Infrastructure Framework</span>
						</h1>

						{/* Subheading */}
						<p className="text-base sm:text-xl lg:text-2xl text-slate-200/90 font-medium leading-relaxed max-w-3xl mx-auto mb-10">
							Author KRM‑style YAML once, validate with Protobuf + Buf ProtoValidate, then execute with Pulumi or OpenTofu.
							Consistent APIs across AWS, GCP, Azure, and Kubernetes—no provider‑specific yak‑shaving.
						</p>

						{/* CTAs */}
						<div className="flex flex-col sm:flex-row gap-4 justify-center items-center mb-12">
							<Button
								size="lg"
								onClick={copyToClipboard}
								className="rounded-full bg-gradient-to-r from-[#7a4183] to-[#d946ef] text-white font-mono text-sm px-8 py-3 h-auto shadow-lg shadow-fuchsia-800/30 transform-gpu transition-all hover:brightness-110 hover:-translate-y-0.5 focus-visible:ring-2 focus-visible:ring-fuchsia-400/50"
							>
								{copied ? (
									<>
										<Check className="w-4 h-4 mr-2" />
										Copied!
									</>
								) : (
									<>
										<Copy className="w-4 h-4 mr-2" />
										{installCommand}
									</>
								)}
							</Button>

							<Button
								size="lg"
								variant="outline"
								onClick={scrollToExample}
								className="rounded-full border border-white/20 text-slate-100 bg-transparent hover:bg-white/5 hover:border-white/40 px-8 py-3 h-auto transform-gpu transition-all hover:-translate-y-0.5 focus-visible:ring-2 focus-visible:ring-fuchsia-400/40"
							>
								Try an Example
								<ArrowDown className="w-4 h-4 ml-2" />
							</Button>
						</div>

						{/* Tertiary link */}
						<div className="flex flex-wrap justify-center gap-6 text-slate-300">
							<a
								href="https://github.com/project-planton/project-planton"
								target="_blank"
								rel="noopener noreferrer"
								className="flex items-center gap-2 hover:text-white transition-colors"
							>
								<Github className="w-4 h-4" />
								View on GitHub
							</a>
						</div>

						{/* Tagline */}
						<div className="mt-16 pt-12 border-t border-white/10">
							<p className="text-2xl sm:text-3xl font-bold text-center">
								<span className="text-slate-300">Define once.</span>
								<span className="ml-3 bg-gradient-to-r from-[#f0abfc] via-[#f472b6] to-[#d946ef] bg-clip-text text-transparent">Deploy anywhere.</span>
							</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}


