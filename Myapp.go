package brainfuck

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func GetRootCmd() *cobra.Command {
	return rootCmd
}

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		// 		Use:   "reconstruct",
		// 		Short: "Reconstruct parses, analyzes, and migrates ASML ASMLMake projects",
		// 		Long: `Reconstruct is a tool that parses, analyses, and migrates ASML ASMLMake projects.

		// Reconstruct extracts information from source code files and (leaf) makefiles.
		// It can use this information to analyze (build) properties of ASML projects/components.
		// In addition, Reconstruct can use and transform this information to automatically
		// migrate to a target build system format, outputting new build files.`,
		// 		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 			// if cpuprofile != "" {
		// 			// 	fCPUprofile, err := os.Create(cpuprofile)
		// 			// 	if err != nil {
		// 			// 		panic(fmt.Sprintf("could not create CPU profile: %v", err))
		// 			// 	}
		// 			// 	if err := pprof.StartCPUProfile(fCPUprofile); err != nil {
		// 			// 		panic(fmt.Sprintf("could not start CPU profile: %v", err))
		// 			// 	}
		// 			// }
		// 		},
		// 		PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// 			// if cpuprofile != "" {
		// 			// 	defer fCPUprofile.Close() // error handling omitted for example
		// 			// 	defer pprof.StopCPUProfile()
		// 			// }
		// 			// if memprofile != "" {
		// 			// 	f, err := os.Create(memprofile)
		// 			// 	if err != nil {
		// 			// 		panic(fmt.Sprintf("could not create memory profile: %v", err))
		// 			// 	}
		// 			// 	defer f.Close() // error handling omitted for example
		// 			// 	runtime.GC()    // get up-to-date statistics
		// 			// 	if err := pprof.WriteHeapProfile(f); err != nil {
		// 			// 		panic(fmt.Sprintf("could not write memory profile: %v", err))
		// 			// 	}
		// 			// }
		// 		},
	}
}

func Initialize(rootCmd1 *cobra.Command) {
	// Create a fresh rootCmd
	rootCmd = rootCmd1

	addIncreaseCommand()
	//addDecCommand()
}

func addIncreaseCommand() {
	// convertCmd represents the convert command
	var convertCmd = &cobra.Command{
		Use:   "+ [scan path]",
		Short: "Convert a project/component to a different build system",
		Long: `Convert a project/component to a different build system.

To scan a specific directory for ASML components, you can either provide the [scan path] directly or use the -s flag.
To force scan a specific ASML component, specify the component directory using the -p flag.

Note: this is still experimental, currently we default to a BOA-derived Bazel conversion.

Reconstruct convert will:
1) Scan and analyse the indicated project/component
2) Resolve external dependencies
3) Postprocess known issues (e.g., resolve cyclic dependencies)
4) Output new build files.`,
		RunE: ProvideLoggedFunc(Run),
	}

	rootCmd.AddCommand(convertCmd)
}

// func addDecCommand() {
// 	// convertCmd represents the convert command
// 	var convertCmd = &cobra.Command{
// 		Use:   "- [scan path]",
// 		Short: "Convert a project/component to a different build system",
// 		Long: `Convert a project/component to a different build system.

// To scan a specific directory for ASML components, you can either provide the [scan path] directly or use the -s flag.
// To force scan a specific ASML component, specify the component directory using the -p flag.

// Note: this is still experimental, currently we default to a BOA-derived Bazel conversion.

// Reconstruct convert will:
// 1) Scan and analyse the indicated project/component
// 2) Resolve external dependencies
// 3) Postprocess known issues (e.g., resolve cyclic dependencies)
// 4) Output new build files.`,
// 		RunE: ProvideLoggedFunc(ExecuteDecrease),
// 	}

// 	rootCmd.AddCommand(convertCmd)
// }

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
