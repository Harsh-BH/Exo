import CodeBlock from '@/components/CodeBlock';

export default function DatabasesOutputPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── GENERATED OUTPUT ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Databases</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-green pl-4">
          Docker Compose configurations for PostgreSQL, MySQL, MongoDB, and Redis.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generation</h2>
        <CodeBlock language="bash" code={`exo gen db --db postgres   # or mysql, mongo, redis`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> PostgreSQL</h2>
        <CodeBlock language="yaml" filename="docker-compose.postgres.yml" showLineNumbers code={`version: '3.8'
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: my-service
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

volumes:
  pgdata:`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> MySQL</h2>
        <CodeBlock language="yaml" filename="docker-compose.mysql.yml" showLineNumbers code={`version: '3.8'
services:
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: secret
      MYSQL_DATABASE: my-service
      MYSQL_USER: admin
      MYSQL_PASSWORD: secret
    ports:
      - "3306:3306"
    volumes:
      - mysqldata:/var/lib/mysql

volumes:
  mysqldata:`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> MongoDB</h2>
        <CodeBlock language="yaml" filename="docker-compose.mongo.yml" showLineNumbers code={`version: '3.8'
services:
  mongo:
    image: mongo:7.0
    environment:
      MONGO_INITDB_ROOT_USERNAME: admin
      MONGO_INITDB_ROOT_PASSWORD: secret
    ports:
      - "27017:27017"
    volumes:
      - mongodata:/data/db

volumes:
  mongodata:`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Redis</h2>
        <CodeBlock language="yaml" filename="docker-compose.redis.yml" showLineNumbers code={`version: '3.8'
services:
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redisdata:/data

volumes:
  redisdata:`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Usage</h2>
        <CodeBlock language="bash" code={`# Start database
docker-compose -f docker-compose.postgres.yml up -d

# Stop database
docker-compose -f docker-compose.postgres.yml down

# Reset (remove data)
docker-compose -f docker-compose.postgres.yml down -v`} />
      </section>
    </article>
  );
}
