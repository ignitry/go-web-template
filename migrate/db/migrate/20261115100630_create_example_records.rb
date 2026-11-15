class CreateExampleRecords < ActiveRecord::Migration[8.0]
  def change
    enable_extension "pgcrypto"

    create_table :example_records, id: :uuid do |t|
      t.string :name, null: false
      t.timestamps
    end
  end
end
