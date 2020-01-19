// Using the go tutorials to build a simple app -- see https://golang.org/doc/articles/wiki/#tmp_13

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(randomize_quote()))
	})
	// https://stackoverflow.com/questions/36751071/heroku-web-process-failed-to-bind-to-port-within-90-seconds-of-launch-tootall
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), r)
	// http.ListenAndServe("0.0.0.0:"+"3000", r)

}

func randomize_quote() string {
	adverb := [32]string{"appropriately", "assertively", "authoritatively", "coherently", "collaboratively", "compellingly", "competently", "completely", "continually", "conveniently", "credibly", "distinctively", "dynamically", "efficiently", "enthusiastically", "globally", "holistically", "hypothetically", "interactively", "intrinsically", "locally", "meaningfully", "monotonically", "objectively", "proactively", "professionally", "progressively", "rapidly", "seamlessly", "synergistically", "totally", "uniquely"}
	verb := [95]string{"actualize", "administrate", "aggregate", "architect", "benchmark", "brand", "build", "communicate", "conceptualize", "coordinate", "create", "cultivate", "customize", "deliver", "deploy", "develop", "disintermediate", "disseminate", "drive", "embrace", "e-enable", "empower", "enable", "engage", "engineer", "enhance", "envision", "evisculate", "evolve", "expedite", "exploit", "extend", "fabricate", "facilitate", "fashion", "formulate", "foster", "generate", "grow", "harness", "impact", "implement", "incentivize", "incubate", "initiate", "innovate", "integrate", "iterate", "leverage", "maintain", "matrix", "maximize", "mesh", "monetize", "morph", "myocardinate", "negotiate", "network", "optimize", "orchestrate", "parallel task", "plagiarize", "pontificate", "predominate", "procrastinate", "productize", "promote", "provide access to", "pursue", "recapitalize", "reconceptualize", "redefine", "re-engineer", "reintermediate", "reinvent", "repurpose", "restore", "revolutionize", "scale", "seize", "simplify", "strategize", "streamline", "supply", "syndicate", "synergize", "synthesize", "target", "transform", "transition", "underwhelm", "unleash", "utilize", "visualize", "whiteboard"}
	adjective := [167]string{"24/7", "accelerated", "accurate", "adaptive", "alternative", "an expanded array of", "antifragile", "B2B", "B2C", "B2G", "B2X", "backend", "backward-compatible", "best-of-breed", "bleeding-edge", "bricks-and-clicks", "bulletproof", "business", "clicks-and-mortar", "client-based", "client-centered", "client-centric", "client-focused", "collaborative", "compelling", "competitive", "complex", "contextualized", "cooperative", "corporate", "cost effective", "covalent", "cradle-to-cradle", "cross functional", "cross-media", "cross-platform", "cross-unit", "customer directed", "customer-driven", "customized", "cutting-edge", "design-for-x", "distinctive", "distributed", "divergent", "diverse", "dynamic", "e-business", "economically sound", "edge of chaos", "effective", "efficient", "emergent", "emerging", "empowered", "enabled", "end-to-end", "enterprise", "enterprise-wide", "equity invested", "error-free", "ethical", "excellent", "exceptional", "extensible", "extensive", "flexible", "focused", "frictionless", "front-end", "fully researched", "fully tested", "functional", "functionalized", "future-proof", "global", "go forward", "goal-oriented", "granular", "high standards in", "high-payoff", "high-quality", "highly efficient", "holistic", "impactful", "inclusive", "inexpensive", "innovative", "installed base", "integrated", "interactive", "interdependent", "intermandated", "interoperable", "intuitive", "just in time", "leading-edge", "leveraged", "long-term high-impact", "low-risk high-yield", "magnetic", "maintainable", "market positioning", "market-driven", "mission-critical", "multidisciplinary", "multifunctional", "multimedia based", "next-generation", "one-to-one", "open-source", "optimal", "orthogonal", "out-of-the-box", "pandemic", "parallel", "performance based", "plug-and-play", "premier", "premium", "principle-centered", "proactive", "process-centric", "professional", "progressive", "prospective", "quality", "real-time", "reliable", "resource sucking", "resource maximizing", "resource-leveling", "revolutionary", "scalable", "seamless", "stand-alone", "standardized", "standards compliant", "state of the art", "sticky", "strategic", "streamlined", "superior", "sustainable", "synergistic", "tactical", "team building", "team driven", "technically sound", "timely", "top-line", "transparent", "turnkey", "ubiquitous", "unique", "user-centric", "user friendly", "value-added", "vertical", "viral", "virtual", "visceral", "visionary", "web-enabled", "wireless", "world-class", "worldwide"}
	noun := [95]string{"action items", "AI", "algorithm", "alignments", "applications", "architectures", "bandwidth", "benefits", "best practices", "blamestorming", "breadboard", "catalysts for change", "channels", "collaboration and idea-sharing", "communities", "content", "convergence", "core competencies", "customer service", "data", "deliverables", "design", "e-business", "e-commerce", "e-markets", "e-tailers", "e-services", "experiences", "expertise", "functionalities", "future", "growth strategies", "human capital", "ideas", "imperatives", "infomediaries", "information", "infrastructures", "initiatives", "innovation", "intellectual capital", "interfaces", "internal or organic sources", "knowledge", "leadership", "leadership skills", "lean", "manufactured products", "markets", "materials", "meaning", "meta-services", "methodologies", "methods of empowerment", "metrics", "mindshare", "models", "networks", "niches", "niche markets", "opportunities", "outside the box thinking", "outsourcing", "paradigms", "partnerships", "pivot", "platforms", "portals", "potentialities", "process improvements", "processes", "products", "quality vectors", "relationships", "resources", "results", "ROI", "scenarios", "schemas", "services", "solutions", "sources", "strategic theme areas", "supply chains", "synergy", "systems", "technologies", "technology", "testing procedures", "total linkage", "users", "value", "vortals", "web-readiness", "web services"}
	// fmt.Println(adverb)
	// fmt.Println(verb)
	// fmt.Println(adjective)
	// fmt.Println(noun)

	rand.Seed(time.Now().Unix())
	innovation_project_gen := fmt.Sprint(adverb[rand.Intn(len(adverb))], " ", verb[rand.Intn(len(verb))], " ", adjective[rand.Intn(len(adjective))], " ", noun[rand.Intn(len(noun))])
	// fmt.Println(innovation_project_gen)
	return innovation_project_gen
}
