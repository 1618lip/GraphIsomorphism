from manim import *
import json
import numpy as np

class GraphIsomorphism(Scene):
    def construct(self):
        # Load isomorphism mapping from JSON
        with open("isomorphism_log.json", "r") as f:
            data = json.load(f)
        mapping = data["mapping"]
        n = len(mapping)
        
        # Create initial positions arranged in a circle for Graph 1.
        graph1_positions = {}
        for i in range(n):
            angle = 2 * PI * i / n
            x = np.cos(angle)
            y = np.sin(angle)
            graph1_positions[i] = np.array([x, y, 0])
        
        # Graph 2 positions: nodes are rearranged according to the mapping.
        # mapping[i] tells where node i of Graph1 should go in Graph2.
        graph2_positions = {}
        for i in range(n):
            # For demonstration, assign the target position as the original position of the node that maps to i.
            # In other words, node mapping[i] in Graph1 goes to position of node i.
            graph2_positions[mapping[i]] = graph1_positions[i]
        
        # Create nodes (circles with labels) for the initial graph layout.
        nodes = VGroup()
        for i in range(n):
            circle = Circle(radius=0.2, color=BLUE).move_to(graph1_positions[i])
            label = Text(str(i), font_size=24).move_to(graph1_positions[i])
            nodes.add(VGroup(circle, label))
        
        self.play(Create(nodes))
        self.wait(1)
        
        # Animate the nodes moving to their positions in Graph 2.
        animations = []
        for i in range(n):
            target_pos = graph2_positions[i]
            animations.append(nodes[i].animate.move_to(target_pos))
        
        self.play(*animations, run_time=2)
        self.wait(1)
        
        # Highlight the final configuration.
        self.play(Indicate(nodes))
        self.wait(2)
