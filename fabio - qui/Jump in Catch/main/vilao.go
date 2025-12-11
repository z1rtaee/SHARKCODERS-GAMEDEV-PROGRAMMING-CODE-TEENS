components {
  id: "enemy"
  component: "/main/assets/scripts/enemy.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Vil\\303\\243o.png\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/tiles/player.atlas\"\n"
  "}\n"
  ""
  scale {
    y: 1.069774
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"vilao\"\n"
  "mask: \"ground\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"box\"\n"
  "  }\n"
  "  data: 12.566866\n"
  "  data: 17.00336\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
