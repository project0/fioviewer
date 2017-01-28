function newColorSet (r, g, b) {
  let rgb = 'rgb(' + r + ',' + g + ',' + b + ')'
  return {
    borderColor: rgb,
    pointHoverBorderColor: rgb,
    backgroundColor: rgb
  }
}

const colors = [
  newColorSet(244, 67, 54),
  newColorSet(233, 30, 99),
  newColorSet(156, 39, 176),
  newColorSet(94, 53, 177),
  newColorSet(57, 73, 171),
  newColorSet(30, 136, 229),
  newColorSet(41, 182, 246),
  newColorSet(0, 172, 193),
  newColorSet(0, 121, 107),
  newColorSet(67, 160, 71),
  newColorSet(139, 195, 74),
  newColorSet(205, 220, 57),
  newColorSet(253, 216, 53),
  newColorSet(255, 179, 0),
  newColorSet(244, 81, 30),
  newColorSet(121, 85, 72),
  newColorSet(96, 125, 139)
]

export default colors
