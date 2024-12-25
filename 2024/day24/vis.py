'''
Thank you chatgpt
'''

from graphviz import Digraph

# Function to parse the input file and extract edges
def parse_edges(file_path):
    edges = []
    with open(file_path, 'r') as f:
        for line in f:
            line = line.strip()
            if line.startswith('-') or line.startswith('<'):  # Skip irrelevant lines
                continue
            if "->" in line:
                inputs, output = line.split("->")
                inputs = inputs.strip()
                output = output.strip()
                edges.append((inputs, output))
    print(f"Parsed edges: {edges}")  # Debug print
    return edges

# Main function to create and visualize the graph
def create_graph(edges):
    dot = Digraph(format='svg')  # Set format to SVG
    dot.attr(rankdir='LR', size='10')  # Left-to-right layout

    if not edges:  # Check if edges are empty
        print("No edges to process. Check the input file.")
        return

    # Process edges to add nodes and connections
    for inputs, output in edges:
        input_parts = inputs.split(" ")
        if len(input_parts) == 3:
            left, op, right = input_parts

            # Create gate node with a distinct style
            gate_node = f"{left}_{right}_{op}"  # Unique ID for the gate
            dot.node(gate_node, op, shape="box", style="filled", fillcolor="lightgrey")

            # Add input and output connections
            dot.node(left, left)  # Input 1
            dot.node(right, right)  # Input 2
            dot.node(output, output, shape="ellipse", style="filled", fillcolor="lightblue")  # Output

            dot.edge(left, gate_node)  # Connect input 1 to gate
            dot.edge(right, gate_node)  # Connect input 2 to gate
            dot.edge(gate_node, output)  # Connect gate to output

    # Save the graph to a file
    output_file = dot.render('logic_network', format='svg', cleanup=True)
    print(f"Graph saved to {output_file}")

# Load edges from file and generate the graph
edges = parse_edges('input.txt')
create_graph(edges)
