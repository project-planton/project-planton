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
			{/* Indigo → Navy gradient */}
			<div className="absolute inset-0 bg-gradient-to-b from-indigo-950 via-indigo-900 to-slate-950" />
			{/* Subtle dot overlay */}
			<div
				className="absolute inset-0 opacity-30"
				style={{
					backgroundImage:
						"url(\"data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23334155' fill-opacity='0.12'%3E%3Ccircle cx='30' cy='30' r='1'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E\")",
				}}
			/>

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
							<span className="bg-gradient-to-r from-white via-fuchsia-300 to-fuchsia-500 bg-clip-text text-transparent">Open‑Source Multi‑Cloud</span>
							<br />
							<span className="bg-gradient-to-r from-white via-fuchsia-300 to-fuchsia-500 bg-clip-text text-transparent">Infrastructure Framework</span>
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
								className="rounded-full bg-[#7a4183] hover:bg-[#7a4183]/90 text-white font-mono text-sm px-8 py-3 h-auto shadow-lg shadow-fuchsia-700/25 transform-gpu transition-all hover:brightness-110 hover:-translate-y-0.5 focus-visible:ring-2 focus-visible:ring-fuchsia-400/50"
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
								className="rounded-full border border-fuchsia-500/40 text-slate-100 bg-white/0 hover:bg-fuchsia-500/10 hover:border-fuchsia-300 px-8 py-3 h-auto transform-gpu transition-all hover:-translate-y-0.5 focus-visible:ring-2 focus-visible:ring-fuchsia-400/40"
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
								<span className="text-[#d946ef] ml-3">Deploy anywhere.</span>
							</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}


