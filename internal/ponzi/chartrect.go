package ponzi

import "github.com/go-gl/gl/v4.5-core/gl"

// createChartBackgroundVAO creates a 2-triangle VAO for a square centered around (0, 0).
func createChartBackgroundVAO(ulColor, urColor, blColor, brColor [3]float32) (uint32, int32) {
	vertices := []float32{
		-1, 1, // UL - 0
		1, 1, // UR - 1
		-1, -1, // BL - 2
		1, -1, // BR - 3
	}

	colors := []float32{
		ulColor[0], ulColor[1], ulColor[2], // UL
		urColor[0], urColor[1], urColor[2], // UR
		blColor[0], blColor[1], blColor[2], // BL
		brColor[0], brColor[1], brColor[2], // BR
	}

	indices := []uint16{
		0, 2, 1,
		1, 2, 3,
	}

	vbo := createArrayBuffer(vertices)
	cbo := createArrayBuffer(colors)
	ibo := createElementArrayBuffer(indices)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	{
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.EnableVertexAttribArray(positionLocation)
		gl.VertexAttribPointer(positionLocation, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))

		gl.BindBuffer(gl.ARRAY_BUFFER, cbo)
		gl.EnableVertexAttribArray(colorLocation)
		gl.VertexAttribPointer(colorLocation, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	}
	gl.BindVertexArray(0)

	return vao, int32(len(indices))
}